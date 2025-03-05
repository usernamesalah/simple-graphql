package openmeteo

type ElevationResponse struct {
	Elevation []float64 `json:"elevation"`
}

type HourlyUnits struct {
	Time             string `json:"time"`
	Temperature2m    string `json:"temperature_2m"`
	Precipitation    string `json:"precipitation"`
	WindSpeed10m     string `json:"wind_speed_10m"`
	WindDirection10m string `json:"wind_direction_10m"`
}

type HourlyData struct {
	Time             []string  `json:"time"`
	Temperature2m    []float64 `json:"temperature_2m"`
	Precipitation    []float64 `json:"precipitation"`
	WindSpeed10m     []float64 `json:"wind_speed_10m"`
	WindDirection10m []float64 `json:"wind_direction_10m"`
}

type WeatherResponse struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationTimeMs     float64     `json:"generationtime_ms"`
	UtcOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               HourlyData  `json:"hourly"`
}
