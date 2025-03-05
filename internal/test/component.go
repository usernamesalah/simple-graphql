package test

import (
	"tensor-graphql/infrastructure/config"
	"tensor-graphql/internal/test/mockrepository"
	"tensor-graphql/internal/test/mockusecase"
	"testing"
)

type MockComponent struct {
	Config               *config.Config
	PowerPlantRepository *mockrepository.PowerPlantRepository
	PowerPlantUsecase    *mockusecase.PowerPlantUsecase
}

func InitMockComponent(t *testing.T) *MockComponent {
	return &MockComponent{
		Config:               &config.Config{},
		PowerPlantRepository: mockrepository.NewPowerPlantRepository(t),
		PowerPlantUsecase:    mockusecase.NewPowerPlantUsecase(t),
	}
}
