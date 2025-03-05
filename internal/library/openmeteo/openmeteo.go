package openmeteo

import (
	"context"
	"encoding/json"
	"fmt"
	"tensor-graphql/pkg/derrors"

	"github.com/go-resty/resty/v2"
)

const (
	openMeteoAPI = "https://api.open-meteo.com/v1/"
)

type (
	OpenMeteo struct {
		api *resty.Client
	}

	openmeteo interface {
		GetWeatherForecast(ctx context.Context, latitude, longitude float64, days int) (weather *WeatherResponse, err error)
	}
)

func NewOpenMeteo() OpenMeteo {
	return OpenMeteo{
		api: resty.New(),
	}
}

// https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m,precipitation,wind_speed_10m,wind_direction_10m
func (o *OpenMeteo) GetWeatherForecast(ctx context.Context, latitude, longitude float64, days int) (weather *WeatherResponse, err error) {
	defer derrors.Wrap(&err, "GetWeatherForecast(%f,%f)", latitude, longitude)

	if days == 0 {
		days = 7
	}

	url := fmt.Sprintf("%sforecast?latitude=%f&longitude=%f&hourly=temperature_2m,precipitation,wind_speed_10m,wind_direction_10m&forecast_days=%d", openMeteoAPI, latitude, longitude, days)
	resp, err := o.api.R().
		Get(url)
	if err != nil {
		return weather, err
	}

	err = json.Unmarshal(resp.Body(), &weather)
	if err != nil {
		return weather, err
	}

	return
}
