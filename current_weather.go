package openmeteogo

import (
	"context"
)

type CurrentWeatherService service

const (
	CurrentTemperature2m       OpenMeteoConst = "temperature_2m"
	CurrentRelativeHumidity2m  OpenMeteoConst = "relative_humidity_2m"
	CurrentApparentTemperature OpenMeteoConst = "apparent_temperature"
	CurrentIsDay               OpenMeteoConst = "is_day"
	CurrentPrecipitation       OpenMeteoConst = "precipitation"
	CurrentRain                OpenMeteoConst = "rain"
	CurrentShowers             OpenMeteoConst = "showers"
	CurrentSnowfall            OpenMeteoConst = "snowfall"
	CurrentCloudCover          OpenMeteoConst = "cloud_cover"
	CurrentWeatherCode         OpenMeteoConst = "weather_code"
	CurrentSealevelPressure    OpenMeteoConst = "pressure_msl"
	CurrentSurfacePressure     OpenMeteoConst = "surface_pressure"
	CurrentWindSpeed10m        OpenMeteoConst = "wind_speed_10m"
	CurrentWindDirection10m    OpenMeteoConst = "wind_direction_10m"
	CurrentWindGust10m         OpenMeteoConst = "wind_gusts_10m"
)

type CurrentOptions struct {
	Options
	Latitude  float64           `url:"latitude"`
	Longitude float64           `url:"longitude"`
	Current   *[]OpenMeteoConst `url:"current,omitempty"`
}

type CurrentWeatherResponse struct {
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Generationtime_ms    float64               `json:"generationtime_ms"`
	UtcOffsetSeconds     int                   `json:"utc_offset_seconds"`
	Timezone             string                `json:"timezone"`
	TimezoneAbbreviation string                `json:"timezone_abbreviation"`
	Elevation            float64               `json:"elevation"`
	CurrentUnits         *CurrentUnitsResponse `json:"current_units"`
	Current              *CurrentResponse      `json:"current"`
}

type CurrentUnitsResponse struct {
	Time                string  `json:"time,omitempty"`
	Interval            string  `json:"interval,omitempty"`
	Temperature2m       *string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  *string `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature *string `json:"apparent_temperature,omitempty"`
	IsDay               *string `json:"is_day,omitempty"`
	Precipitation       *string `json:"precipitation,omitempty"`
	Rain                *string `json:"rain,omitempty"`
	Showers             *string `json:"showers,omitempty"`
	Snowfall            *string `json:"snowfall,omitempty"`
	WeatherCode         *string `json:"weather_code,omitempty"`
	CloudCover          *string `json:"cloud_cover,omitempty"`
	SealevelPressure    *string `json:"pressure_msl,omitempty"`
	SurfacePressure     *string `json:"surface_pressure,omitempty"`
	WindSpeed10m        *string `json:"wind_speed_10m,omitempty"`
	WindDirection10m    *string `json:"wind_direction_10m,omitempty"`
	WindGust10m         *string `json:"wind_gusts_10m,omitempty"`
}

type CurrentResponse struct {
	Time                string   `json:"time,omitempty"`
	Interval            int      `json:"interval,omitempty"`
	Temperature2m       *float64 `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  *float64 `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature *float64 `json:"apparent_temperature,omitempty"`
	IsDay               *int     `json:"is_day,omitempty"`
	Precipitation       *float64 `json:"precipitation,omitempty"`
	Rain                *float64 `json:"rain,omitempty"`
	Showers             *float64 `json:"showers,omitempty"`
	Snowfall            *float64 `json:"snowfall,omitempty"`
	WeatherCode         *int     `json:"weather_code,omitempty"`
	CloudCover          *float64 `json:"cloud_cover,omitempty"`
	SealevelPressure    *float64 `json:"pressure_msl,omitempty"`
	SurfacePressure     *float64 `json:"surface_pressure,omitempty"`
	WindSpeed10m        *float64 `json:"wind_speed_10m,omitempty"`
	WindDirection10m    *int     `json:"wind_direction_10m,omitempty"`
	WindGust10m         *float64 `json:"wind_gusts_10m,omitempty"`
}

// Forecast retrieve current weather based on provide location and requested params
// Forecast days not be allowed in this request
func (service *CurrentWeatherService) Forecast(ctx context.Context, opts *CurrentOptions) (*CurrentWeatherResponse, error) {
	u, err := addOptions("forecast/", opts)

	if err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("GET", service.client.WeatherBaseURL, u, nil)

	if err != nil {
		return nil, err
	}

	result := new(CurrentWeatherResponse)

	_, err = service.client.Do(ctx, req, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
