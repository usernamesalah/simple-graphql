package config

import (
	"fmt"
	"log"
	"tensor-graphql/internal/constant"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment string
	HttpPort    string
	CORSOrigins []string

	DBMaster *DB
	DBSlave  *DB
}

// DB config model
type DB struct {
	ConnectionString string
	MaxIdle          int
	MaxOpen          int
}

// DatabaseConfig stores database configurations.
type configEnv struct {
	Port        string   `envconfig:"APP_PORT" default:"8080"`
	Env         string   `envconfig:"APP_ENV" default:"development"`
	CORSOrigins []string `envconfig:"CORS_ORIGINS"`

	// Database config
	DBMasterMaxIdle int    `envconfig:"DBMASTERMAXIDLECONN"`
	DBMasterMaxOpen int    `envconfig:"DBMASTERMAXOPENCONN"`
	DBSlaveMaxIdle  int    `envconfig:"DBSLAVEMAXIDLECONN"`
	DBSlaveMaxOpen  int    `envconfig:"DBSLAVEMAXOPENCONN"`
	DBMasterUser    string `envconfig:"DBMASTERUSER"`
	DBMasterPass    string `envconfig:"DBMASTERPASS"`
	DBMasterHost    string `envconfig:"DBMASTERHOST"`
	DBMasterPort    string `envconfig:"DBMASTERPORT"`
	DBMasterName    string `envconfig:"DBMASTERNAME"`
	DBSlaveUser     string `envconfig:"DBSLAVEUSER"`
	DBSlavePass     string `envconfig:"DBSLAVEPASS"`
	DBSlaveHost     string `envconfig:"DBSLAVEHOST"`
	DBSlavePort     string `envconfig:"DBSLAVEPORT"`
	DBSlaveName     string `envconfig:"DBSLAVENAME"`
}

var appConfig *Config

// ReadConfig populates configurations from environment variables.
func Init() {
	_ = godotenv.Overload()
	var cfg configEnv
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("[Init] failed to map config, %+v\n", err)
	}

	appConfig = &Config{}
	appConfig.Environment = cfg.Env
	appConfig.HttpPort = cfg.Port
	appConfig.CORSOrigins = cfg.CORSOrigins

	initDB(&cfg)
}

func initDB(c *configEnv) {
	appConfig.DBMaster = &DB{
		ConnectionString: fmt.Sprintf(
			constant.DBStringConnection,
			c.DBMasterUser,
			c.DBMasterPass,
			c.DBMasterHost,
			c.DBMasterPort,
			c.DBMasterName,
		),
		MaxIdle: c.DBMasterMaxIdle,
		MaxOpen: c.DBMasterMaxOpen,
	}
	appConfig.DBSlave = &DB{
		ConnectionString: fmt.Sprintf(
			constant.DBStringConnection,
			c.DBSlaveUser,
			c.DBSlavePass,
			c.DBSlaveHost,
			c.DBSlavePort,
			c.DBSlaveName,
		),
		MaxIdle: c.DBMasterMaxIdle,
		MaxOpen: c.DBMasterMaxOpen,
	}
}

// Get private instance config
func Get() *Config {
	return appConfig
}
