package graphql

import (
	"tensor-graphql/internal/library/openmeteo"
	"tensor-graphql/internal/model"
	usecase "tensor-graphql/internal/usecase/power_plant"
)

// Resolver adalah root resolver yang menyimpan dependency usecase.
type Resolver struct {
	PowerPlantUsecase usecase.PowerPlantUsecase
	OpenmeteoLib      openmeteo.OpenMeteo
}

func NewResolver(powerplantUsecase usecase.PowerPlantUsecase, openmeteoLib openmeteo.OpenMeteo) *Resolver {
	return &Resolver{
		PowerPlantUsecase: powerplantUsecase,
		OpenmeteoLib:      openmeteoLib,
	}
}

func mapToModel(plant *model.PowerPlant, weather *openmeteo.WeatherResponse) (*model.PowerPlant, error) {
	mp := &model.PowerPlant{
		ID:        plant.ID,
		Name:      plant.Name,
		Latitude:  plant.Latitude,
		Longitude: plant.Longitude,
		Elevation: weather.Elevation,
	}

	forecastDays := 7
	if len(weather.Hourly.Time) < forecastDays {
		forecastDays = len(weather.Hourly.Time)
	}

	var forecasts []*model.WeatherForecast
	for i := 0; i < forecastDays; i++ {
		wf := &model.WeatherForecast{
			Time:          weather.Hourly.Time[i],
			Temperature:   weather.Hourly.Temperature2m[i],
			Precipitation: weather.Hourly.Precipitation[i],
			WindSpeed:     weather.Hourly.WindSpeed10m[i],
			WindDirection: weather.Hourly.WindDirection10m[i],
		}
		forecasts = append(forecasts, wf)
	}
	mp.WeatherForecasts = forecasts

	if len(weather.Hourly.Precipitation) > 0 && weather.Hourly.Precipitation[0] > 0 {
		mp.HasPrecipitationToday = true
	} else {
		mp.HasPrecipitationToday = false
	}

	return mp, nil
}
