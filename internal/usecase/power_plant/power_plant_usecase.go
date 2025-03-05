package powerplantusecase

import (
	"context"
	"tensor-graphql/internal/model"
	powerplantrepo "tensor-graphql/internal/repository/power_plant"
	"tensor-graphql/pkg/derrors"
)

type (
	PowerPlantUsecase interface {
		CreatePowerPlant(ctx context.Context, powerplant *model.PowerPlant) (err error)
		GetPowerPlantByID(ctx context.Context, powerplantID string) (powerplant *model.PowerPlant, err error)
		GetPowerPlants(ctx context.Context, page, limit int) (powerplants []*model.PowerPlant, total int, err error)
		UpdatePowerPlant(ctx context.Context, powerplant *model.PowerPlant) (err error)
	}

	powerplantUsecase struct {
		powerplantRepo powerplantrepo.PowerPlantRepository
	}
)

func NewPowerPlantUsecase(powerplantRepo powerplantrepo.PowerPlantRepository) PowerPlantUsecase {
	return &powerplantUsecase{
		powerplantRepo: powerplantRepo,
	}
}

func (u *powerplantUsecase) CreatePowerPlant(ctx context.Context, powerplant *model.PowerPlant) (err error) {
	defer derrors.Wrap(&err, "CreatePowerPlant(%q)", powerplant.Name)

	err = u.powerplantRepo.CreatePowerPlant(ctx, nil, powerplant)

	return
}

func (u *powerplantUsecase) UpdatePowerPlant(ctx context.Context, powerplant *model.PowerPlant) (err error) {
	defer derrors.Wrap(&err, "UpdatePowerPlant(%q)", powerplant.ID)

	err = u.powerplantRepo.UpdatePowerPlant(ctx, nil, powerplant)
	return
}

func (u *powerplantUsecase) GetPowerPlants(ctx context.Context, page, limit int) (powerplants []*model.PowerPlant, total int, err error) {
	defer derrors.Wrap(&err, "CreatePowerPlant")

	powerplants, total, err = u.powerplantRepo.GetPowerPlants(ctx, page, limit)
	return
}

func (u *powerplantUsecase) GetPowerPlantByID(ctx context.Context, powerplantID string) (powerplant *model.PowerPlant, err error) {
	defer derrors.Wrap(&err, "GetPowerPlantByID(%q)", powerplantID)

	powerplant, err = u.powerplantRepo.GetPowerPlantByID(ctx, powerplantID)
	return
}
