package openmeteogo

import (
	"context"
)

type DailyWeatherService service

const (
	DailyWeatherCode                 OpenMeteoConst = "weather_code"
	DailyTemperature2mMax            OpenMeteoConst = "temperature_2m_max"
	DailyTemperature2mMin            OpenMeteoConst = "temperature_2m_min"
	DailyApparentTemperatureMax      OpenMeteoConst = "apparent_temperature_max"
	DailyApparentTemperatureMin      OpenMeteoConst = "apparent_temperature_min"
	DailySunrise                     OpenMeteoConst = "sunrise"
	DailySunset                      OpenMeteoConst = "sunset"
	DailyDaylightDuration            OpenMeteoConst = "daylight_duration"
	DailySunshineDuration            OpenMeteoConst = "sunshine_duration"
	DailyUvIndexMax                  OpenMeteoConst = "uv_index_max"
	DailyUvIndexClearSkyMax          OpenMeteoConst = "uv_index_clear_sky_max"
	DailyPrecipitationSum            OpenMeteoConst = "precipitation_sum"
	DailyRainSum                     OpenMeteoConst = "rain_sum"
	DailyShowersSum                  OpenMeteoConst = "showers_sum"
	DailySnowfallSum                 OpenMeteoConst = "snowfall_sum"
	DailyPrecipitationHours          OpenMeteoConst = "precipitation_hours"
	DailyPrecipitationProbabilityMax OpenMeteoConst = "precipitation_probability_max"
	DailyWindSpeed10mMax             OpenMeteoConst = "wind_speed_10m_max"
	DailyWindGusts10mMax             OpenMeteoConst = "wind_gusts_10m_max"
	DailyWindDirection10mDominant    OpenMeteoConst = "wind_direction_10m_dominant"
	DailyShortwaveRadiationSum       OpenMeteoConst = "shortwave_radiation_sum"
	DailyEt0FaoEvapotranspiration    OpenMeteoConst = "et0_fao_evapotranspiration"
)

type DailyOptions struct {
	Options
	Latitude     float64           `url:"latitude"`
	Longitude    float64           `url:"longitude"`
	Daily        *[]OpenMeteoConst `url:"daily"`
	ForecastDays int               `url:"forecast_days"`
	PastDays     int               `url:"past_days"`
}

type DailyWeatherResponse struct {
	Latitude             float64            `json:"latitude"`
	Longitude            float64            `json:"longitude"`
	Generationtime_ms    float64            `json:"generationtime_ms"`
	UtcOffsetSeconds     int                `json:"utc_offset_seconds"`
	Timezone             string             `json:"timezone"`
	TimezoneAbbreviation string             `json:"timezone_abbreviation"`
	Elevation            float64            `json:"elevation"`
	DailyUnits           DailyUnitsResponse `json:"daily_units"`
	Daily                DailyResponse      `json:"daily"`
}

type DailyUnitsResponse struct {
	Time                        string  `json:"time,omitempty"`
	WeatherCode                 *string `json:"weather_code,omitempty"`
	Temperature2mMax            *string `json:"temperature_2m_max,omitempty"`
	Temperature2mMin            *string `json:"temperature_2m_min,omitempty"`
	ApparentTemperatureMax      *string `json:"apparent_temperature_max,omitempty"`
	ApparentTemperatureMin      *string `json:"apparent_temperature_min,omitempty"`
	Sunrise                     *string `json:"sunrise,omitempty"`
	Sunset                      *string `json:"sunset,omitempty"`
	DaylightDuration            *string `json:"daylight_duration,omitempty"`
	SunshineDuration            *string `json:"sunshine_duration,omitempty"`
	UvIndexMax                  *string `json:"uv_index_max,omitempty"`
	UvIndexClearSkyMax          *string `json:"uv_index_clear_sky_max,omitempty"`
	PrecipitationSum            *string `json:"precipitation_sum,omitempty"`
	RainSum                     *string `json:"rain_sum,omitempty"`
	ShowersSum                  *string `json:"showers_sum,omitempty"`
	SnowfallSum                 *string `json:"snowfall_sum,omitempty"`
	PrecipitationHours          *string `json:"precipitation_hours,omitempty"`
	PrecipitationProbabilityMax *string `json:"precipitation_probability_max,omitempty"`
	WindSpeed10mMax             *string `json:"wind_speed_10m_max,omitempty"`
	WindGusts10mMax             *string `json:"wind_gusts_10m_max,omitempty"`
	WindDirection10mDominant    *string `json:"wind_direction_10m_dominant,omitempty"`
	ShortwaveRadiationSum       *string `json:"shortwave_radiation_sum,omitempty"`
	Et0FaoEvapotranspiration    *string `json:"et0_fao_evapotranspiration,omitempty"`
}

type DailyResponse struct {
	Time                        []string   `json:"time,omitempty"`
	WeatherCode                 []*int32   `json:"weather_code,omitempty"`
	Temperature2mMax            []*float64 `json:"temperature_2m_max,omitempty"`
	Temperature2mMin            []*float64 `json:"temperature_2m_min,omitempty"`
	ApparentTemperatureMax      []*float64 `json:"apparent_temperature_max,omitempty"`
	ApparentTemperatureMin      []*float64 `json:"apparent_temperature_min,omitempty"`
	Sunrise                     []*string  `json:"sunrise,omitempty"`
	Sunset                      []*string  `json:"sunset,omitempty"`
	DaylightDuration            []*float64 `json:"daylight_duration,omitempty"`
	SunshineDuration            []*float64 `json:"sunshine_duration,omitempty"`
	UvIndexMax                  []*float64 `json:"uv_index_max,omitempty"`
	UvIndexClearSkyMax          []*float64 `json:"uv_index_clear_sky_max,omitempty"`
	PrecipitationSum            []*float64 `json:"precipitation_sum,omitempty"`
	RainSum                     []*float64 `json:"rain_sum,omitempty"`
	ShowersSum                  []*float64 `json:"showers_sum,omitempty"`
	SnowfallSum                 []*float64 `json:"snowfall_sum,omitempty"`
	PrecipitationHours          []*float64 `json:"precipitation_hours,omitempty"`
	PrecipitationProbabilityMax []*float64 `json:"precipitation_probability_max,omitempty"`
	WindSpeed10mMax             []*float64 `json:"wind_speed_10m_max,omitempty"`
	WindGusts10mMax             []*float64 `json:"wind_gusts_10m_max,omitempty"`
	WindDirection10mDominant    []*int32   `json:"wind_direction_10m_dominant,omitempty"`
	ShortwaveRadiationSum       []*float64 `json:"shortwave_radiation_sum,omitempty"`
	Et0FaoEvapotranspiration    []*float64 `json:"et0_fao_evapotranspiration,omitempty"`
}

// Forecast retrieve daily weather based on provide location and requested params
func (service *DailyWeatherService) Forecast(ctx context.Context, opts *DailyOptions) (*DailyWeatherResponse, error) {
	u, err := addOptions("forecast/", opts)

	if err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	result := new(DailyWeatherResponse)

	_, err = service.client.Do(ctx, req, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
