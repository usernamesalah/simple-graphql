package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tensor-graphql/infrastructure/config"
	"tensor-graphql/infrastructure/database"
	"tensor-graphql/internal/container"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	graphqlResolver "tensor-graphql/internal/api/graphql"
)

// @title Api Documentation for GraphQL API
// @version 0.1
// @description GraphQL API for managing power plants with weather and geography data.
// @contact.name Tensor Energy
// @contact.email no-reply@tensorenergy.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

func main() {
	// Load configuration
	config.Init()

	// Setup logging
	configLog := zap.NewProductionConfig()
	configLog.EncoderConfig.StacktraceKey = "" // Hide stacktrace info
	configLog.DisableCaller = true

	log, err := configLog.Build()
	if err != nil {
		panic(err)
	}

	if err := run(log); err != nil {
		log.Error("error: shutting down", zap.Error(err))
		os.Exit(1)
	}
}

func run(log *zap.Logger) error {
	conf := config.Get()

	// Initialize Database
	db, err := database.InitializeDatabase(conf)
	if err != nil {
		log.Error("failed to initialize db", zap.Error(err))
		return err
	}

	// Shared component for dependency injection
	sharedComponent := &container.SharedComponent{
		DB:   db,
		Conf: conf,
		Log:  log,
	}

	// Initialize handler components, including GraphQL resolver and use cases
	cc := container.NewHandlerComponent(sharedComponent)

	// Initialize Echo Web Server
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Set("Cache-Control", "max-age=3600, public")
			return next(c)
		}
	})
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info("request",
				zap.String("Latency", v.Latency.String()),
				zap.String("Remote IP", c.RealIP()),
				zap.String("URI", v.URI),
				zap.String("Method", c.Request().Method),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	e.Validator = &requestValidator{}

	// Set up GraphQL handler with transports (POST and GET)
	graphqlHandler := handler.New(graphqlResolver.NewExecutableSchema(graphqlResolver.Config{Resolvers: cc.Resolver}))
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.GET{})

	// GraphQL endpoint
	e.POST("/graphql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	fmt.Println(conf.Environment)
	// Enable Playground only in development environment
	if conf.Environment == "development" {
		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	// Configure HTTP server
	server := &http.Server{
		Addr:         "0.0.0.0:" + conf.HttpPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	serverErrors := make(chan error, 1)
	// Start the server in a goroutine
	go func() {
		log.Info("server listening on", zap.String("address", server.Addr))
		serverErrors <- e.StartServer(server)
	}()

	// Capture shutdown signals for graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("starting server: %v", err)
	case <-shutdown:
		log.Info("caught signal, shutting down")
		const timeout = 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Use Echo's Shutdown to gracefully stop the server and its components
		if err := e.Shutdown(ctx); err != nil {
			log.Error("error gracefully shutting down server", zap.Error(err))
			return fmt.Errorf("could not stop server gracefully: %v", err)
		}
	}

	return nil
}

// requestValidator uses govalidator to validate incoming requests.
type requestValidator struct{}

func (rv *requestValidator) Validate(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	return err
}
