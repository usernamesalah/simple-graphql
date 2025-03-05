package container

import (
	"tensor-graphql/infrastructure/config"
	"tensor-graphql/internal/api/graphql"
	"tensor-graphql/internal/library/openmeteo"
	repository "tensor-graphql/internal/repository/common"
	powerPlantrepository "tensor-graphql/internal/repository/power_plant"
	powerplantusecase "tensor-graphql/internal/usecase/power_plant"
)

type HandlerComponent struct {
	Config   *config.Config
	Resolver *graphql.Resolver

	// UseCase
	PowerPlantUsecase powerplantusecase.PowerPlantUsecase
}

func NewHandlerComponent(sc *SharedComponent) *HandlerComponent {

	baseStore := repository.NewRepository(sc.DB)
	openmeteoLib := openmeteo.NewOpenMeteo()

	powerPlantrepository := powerPlantrepository.NewPowerPlantRepository(baseStore)
	powerplantUsecase := powerplantusecase.NewPowerPlantUsecase(powerPlantrepository)

	resolver := graphql.NewResolver(powerplantUsecase, openmeteoLib)

	return &HandlerComponent{
		Config:   sc.Conf,
		Resolver: resolver,

		PowerPlantUsecase: powerplantUsecase,
	}
}
