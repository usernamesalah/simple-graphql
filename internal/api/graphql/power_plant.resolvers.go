package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"tensor-graphql/internal/model"
)

// CreatePowerPlant is the resolver for the createPowerPlant field.
func (r *mutationResolver) CreatePowerPlant(ctx context.Context, name string, latitude float64, longitude float64) (*model.PowerPlant, error) {
	plant := &model.PowerPlant{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}

	err := r.PowerPlantUsecase.CreatePowerPlant(ctx, plant)
	if err != nil {
		return nil, err
	}

	weather, err := r.OpenmeteoLib.GetWeatherForecast(ctx, plant.Latitude, plant.Longitude, 7)
	if err != nil {
		return nil, err
	}

	return mapToModel(plant, weather)
}

// UpdatePowerPlant is the resolver for the updatePowerPlant field.
func (r *mutationResolver) UpdatePowerPlant(ctx context.Context, id string, name *string, latitude *float64, longitude *float64) (*model.PowerPlant, error) {
	plant := &model.PowerPlant{
		ID:        id,
		Name:      *name,
		Latitude:  *latitude,
		Longitude: *longitude,
	}

	err := r.PowerPlantUsecase.UpdatePowerPlant(ctx, plant)
	if err != nil {
		return nil, err
	}

	weather, err := r.OpenmeteoLib.GetWeatherForecast(ctx, plant.Latitude, plant.Longitude, 7)
	if err != nil {
		return nil, err
	}

	return mapToModel(plant, weather)
}

// PowerPlant is the resolver for the powerPlant field.
func (r *queryResolver) PowerPlant(ctx context.Context, id string) (*model.PowerPlant, error) {
	plant, err := r.PowerPlantUsecase.GetPowerPlantByID(ctx, id)
	if err != nil {
		return nil, err
	}

	weather, err := r.OpenmeteoLib.GetWeatherForecast(ctx, plant.Latitude, plant.Longitude, 7)
	if err != nil {
		return nil, err
	}

	return mapToModel(plant, weather)
}

// PowerPlants is the resolver for the powerPlants field.
func (r *queryResolver) PowerPlants(ctx context.Context, page *int, pageSize *int) (*model.PowerPlantPage, error) {
	if page == nil {
		page = new(int)
		*page = 1
	}
	if pageSize == nil {
		pageSize = new(int)
		*pageSize = 10
	}

	plants, total, err := r.PowerPlantUsecase.GetPowerPlants(ctx, *page, *pageSize)
	if err != nil {
		return nil, err
	}

	var modelPlants []*model.PowerPlant
	for _, plant := range plants {
		weather, err := r.OpenmeteoLib.GetWeatherForecast(ctx, plant.Latitude, plant.Longitude, 7)
		if err != nil {
			return nil, err
		}
		mapped, err := mapToModel(plant, weather)
		if err != nil {
			return nil, err
		}
		modelPlants = append(modelPlants, mapped)
	}

	return &model.PowerPlantPage{
		Plants:     modelPlants,
		TotalCount: total,
		Page:       *page,
		PageSize:   *pageSize,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
