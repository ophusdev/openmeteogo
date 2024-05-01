package openmeteogo

import (
	"context"
)

type HourlyWeatherService service

const (
	HourlyTemperature2m            OpenMeteoConst = "temperature_2m"
	HourlyRelativeHumidity2m       OpenMeteoConst = "relative_humidity_2m"
	HourlyDewpoint2m               OpenMeteoConst = "dew_point_2m"
	HourlyApparentTemperature      OpenMeteoConst = "apparent_temperature"
	HourlyPrecipitationProbability OpenMeteoConst = "precipitation_probability"
	HourlyPrecipitation            OpenMeteoConst = "precipitation"
	HourlyRain                     OpenMeteoConst = "rain"
	HourlyShowers                  OpenMeteoConst = "showers"
	HourlySnowfall                 OpenMeteoConst = "snowfall"
	HourlySnowDepth                OpenMeteoConst = "snow_depth"
	HourlyWeathercode              OpenMeteoConst = "weather_code"
	HourlySealevelPressure         OpenMeteoConst = "pressure_msl"
	HourlySurfacePressure          OpenMeteoConst = "surface_pressure"
	HourlyCloudcoverTotal          OpenMeteoConst = "cloud_cover"
	HourlyCloudcoverLow            OpenMeteoConst = "cloud_cover_low"
	HourlyCloudcoverMid            OpenMeteoConst = "cloud_cover_mid"
	HourlyCloudcoverHigh           OpenMeteoConst = "cloud_cover_high"
	HourlyVisibility               OpenMeteoConst = "visibility"
	HourlyEvapotranspiration       OpenMeteoConst = "evapotranspiration"
	HourlyEt0FaoEvapotranspiration OpenMeteoConst = "et0_fao_evapotranspiration"
	HourlyVapourPressureDeficit    OpenMeteoConst = "vapour_pressure_deficit"
	HourlyWindSpeed10m             OpenMeteoConst = "wind_speed_10m"
	HourlyWindSpeed80m             OpenMeteoConst = "wind_speed_80m"
	HourlyWindSpeed120m            OpenMeteoConst = "wind_speed_120m"
	HourlyWindSpeed180m            OpenMeteoConst = "wind_speed_180m"
	HourlyWindDirection10m         OpenMeteoConst = "wind_direction_10m"
	HourlyWindDirection80m         OpenMeteoConst = "wind_direction_80m"
	HourlyWindDirection120m        OpenMeteoConst = "wind_direction_120m"
	HourlyWindDirection180m        OpenMeteoConst = "wind_direction_180m"
	HourlyWindGusts10m             OpenMeteoConst = "wind_gusts_10m"
	HourlyTemperature80m           OpenMeteoConst = "temperature_80m"
	HourlyTemperature120m          OpenMeteoConst = "temperature_120m"
	HourlyTemperature180m          OpenMeteoConst = "temperature_180m"
	HourlySoilTemperature0cm       OpenMeteoConst = "soil_temperature_0cm"
	HourlySoilTemperature6cm       OpenMeteoConst = "soil_temperature_6cm"
	HourlySoilTemperature18cm      OpenMeteoConst = "soil_temperature_18cm"
	HourlySoilTemperature54cm      OpenMeteoConst = "soil_temperature_54cm"
	HourlySoilMoisture0to1cm       OpenMeteoConst = "soil_moisture_0_to_1cm"
	HourlySoilMoisture1to3cm       OpenMeteoConst = "soil_moisture_1_to_3cm"
	HourlySoilMoisture3to9cm       OpenMeteoConst = "soil_moisture_3_to_9cm"
	HourlySoilMoisture9to27cm      OpenMeteoConst = "soil_moisture_9_to_27cm"
	HourlySoilMoisture27to81cm     OpenMeteoConst = "soil_moisture_27_to_81cm"
	HourlyUvIndex                  OpenMeteoConst = "uv_index"
	HourlyUvIndexClearSky          OpenMeteoConst = "uv_index_clear_sky"
	HourlyIsDayOrNight             OpenMeteoConst = "is_day"
	HourlyCape                     OpenMeteoConst = "cape"
	HourlyFreezingLevelHeight      OpenMeteoConst = "freezing_level_height"
	HourlySunshineDuration         OpenMeteoConst = "sunshine_duration"
)

type HourlyOptions struct {
	Options
	Latitude     float64           `url:"latitude"`
	Longitude    float64           `url:"longitude"`
	Hourly       *[]OpenMeteoConst `url:"hourly"`
	ForecastDays int               `url:"forecast_days"`
	PastDays     int               `url:"past_days"`
}

type HourlyWeatherResponse struct {
	Latitude             float64             `json:"latitude"`
	Longitude            float64             `json:"longitude"`
	Generationtime_ms    float64             `json:"generationtime_ms"`
	UtcOffsetSeconds     int                 `json:"utc_offset_seconds"`
	Timezone             string              `json:"timezone"`
	TimezoneAbbreviation string              `json:"timezone_abbreviation"`
	Elevation            float64             `json:"elevation"`
	HourlyUnits          HourlyUnitsResponse `json:"hourly_units"`
	Hourly               HourlyResponse      `json:"hourly"`
}

type HourlyUnitsResponse struct {
	Time                     string  `json:"time"`
	Temperature2m            *string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       *string `json:"relative_humidity_2m,omitempty"`
	Dewpoint2m               *string `json:"dew_point_2m,omitempty"`
	ApparentTemperature      *string `json:"apparent_temperature,omitempty"`
	PrecipitationProbability *string `json:"precipitation_probability,omitempty"`
	Precipitation            *string `json:"precipitation,omitempty"`
	Rain                     *string `json:"rain,omitempty"`
	Showers                  *string `json:"showers,omitempty"`
	Snowfall                 *string `json:"snowfall,omitempty"`
	SnowDepth                *string `json:"snow_depth,omitempty"`
	WeatherCode              *string `json:"weather_code,omitempty"`
	SealevelPressure         *string `json:"pressure_msl,omitempty"`
	SurfacePressure          *string `json:"surface_pressure,omitempty"`
	CloudcoverTotal          *string `json:"cloud_cover,omitempty"`
	CloudcoverLow            *string `json:"cloud_cover_low,omitempty"`
	CloudcoverMid            *string `json:"cloud_cover_mid,omitempty"`
	CloudcoverHigh           *string `json:"cloud_cover_high,omitempty"`
	Visibility               *string `json:"visibility,omitempty"`
	Evapotranspiration       *string `json:"evapotranspiration,omitempty"`
	Et0FaoEvapotranspiration *string `json:"et0_fao_evapotranspiration,omitempty"`
	VapourPressureDeficit    *string `json:"vapour_pressure_deficit,omitempty"`
	WindSpeed10m             *string `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             *string `json:"wind_speed_80m,omitempty"`
	WindSpeed120m            *string `json:"wind_speed_120m,omitempty"`
	WindSpeed180m            *string `json:"wind_speed_180m,omitempty"`
	WindDirection10m         *string `json:"wind_direction_10m,omitempty"`
	WindDirection80m         *string `json:"wind_direction_80m,omitempty"`
	WindDirection120m        *string `json:"wind_direction_120m,omitempty"`
	WindDirection180m        *string `json:"wind_direction_180m,omitempty"`
	WindGusts10m             *string `json:"wind_gusts_10m,omitempty"`
	Temperature80m           *string `json:"temperature_80m,omitempty"`
	Temperature120m          *string `json:"temperature_120m,omitempty"`
	Temperature180m          *string `json:"temperature_180m,omitempty"`
	SoilTemperature0cm       *string `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm       *string `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm      *string `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm      *string `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture0to1cm       *string `json:"soil_moisture_0_to_1cm,omitempty"`
	SoilMoisture1to3cm       *string `json:"soil_moisture_1_to_3cm,omitempty"`
	SoilMoisture3to9m        *string `json:"soil_moisture_3_to_9cm,omitempty"`
	SoilMoisture9to27cm      *string `json:"soil_moisture_9_to_27cm,omitempty"`
	SoilMoisture27to81cm     *string `json:"soil_moisture_27_to_81cm,omitempty"`
	UvIndex                  *string `json:"uv_index"`
	UvIndexClearSky          *string `json:"uv_index_clear_sky"`
	IsDayOrNight             *string `json:"is_day"`
	Cape                     *string `json:"cape"`
	FreezingLevelHeight      *string `json:"freezing_level_height"`
	SunshineDuration         *string `json:"sunshine_duration"`
}

type HourlyResponse struct {
	Time                     []string   `json:"time"`
	Temperature2m            []*float64 `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       []*int     `json:"relative_humidity_2m,omitempty"`
	Dewpoint2m               []*float64 `json:"dew_point_2m,omitempty"`
	ApparentTemperature      []*float64 `json:"apparent_temperature,omitempty"`
	PrecipitationProbability []*float64 `json:"precipitation_probability,omitempty"`
	Precipitation            []*float64 `json:"precipitation,omitempty"`
	Rain                     []*float64 `json:"rain,omitempty"`
	Showers                  []*float64 `json:"showers,omitempty"`
	Snowfall                 []*float64 `json:"snowfall,omitempty"`
	SnowDepth                []*float64 `json:"snow_depth,omitempty"`
	WeatherCode              []*float64 `json:"weather_code,omitempty"`
	SealevelPressure         []*float64 `json:"pressure_msl,omitempty"`
	SurfacePressure          []*float64 `json:"surface_pressure,omitempty"`
	CloudcoverTotal          []*float64 `json:"cloud_cover,omitempty"`
	CloudcoverLow            []*float64 `json:"cloud_cover_low,omitempty"`
	CloudcoverMid            []*float64 `json:"cloud_cover_mid,omitempty"`
	CloudcoverHigh           []*float64 `json:"cloud_cover_high,omitempty"`
	Visibility               []*float64 `json:"visibility,omitempty"`
	Evapotranspiration       []*float64 `json:"evapotranspiration,omitempty"`
	Et0FaoEvapotranspiration []*float64 `json:"et0_fao_evapotranspiration,omitempty"`
	VapourPressureDeficit    []*float64 `json:"vapour_pressure_deficit,omitempty"`
	WindSpeed10m             []*float64 `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             []*float64 `json:"wind_speed_80m,omitempty"`
	WindSpeed120m            []*float64 `json:"wind_speed_120m,omitempty"`
	WindSpeed180m            []*float64 `json:"wind_speed_180m,omitempty"`
	WindDirection10m         []*float64 `json:"wind_direction_10m,omitempty"`
	WindDirection80m         []*float64 `json:"wind_direction_80m,omitempty"`
	WindDirection120m        []*float64 `json:"wind_direction_120m,omitempty"`
	WindDirection180m        []*float64 `json:"wind_direction_180m,omitempty"`
	WindGusts10m             []*float64 `json:"wind_gusts_10m,omitempty"`
	Temperature80m           []*float64 `json:"temperature_80m,omitempty"`
	Temperature120m          []*float64 `json:"temperature_120m,omitempty"`
	Temperature180m          []*float64 `json:"temperature_180m,omitempty"`
	SoilTemperature0cm       []*float64 `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm       []*float64 `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm      []*float64 `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm      []*float64 `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture0to1cm       []*float64 `json:"soil_moisture_0_to_1cm,omitempty"`
	SoilMoisture1to3cm       []*float64 `json:"soil_moisture_1_to_3cm,omitempty"`
	SoilMoisture3to9cm       []*float64 `json:"soil_moisture_3_to_9cm,omitempty"`
	SoilMoisture9to27cm      []*float64 `json:"soil_moisture_9_to_27cm,omitempty"`
	SoilMoisture27to81cm     []*float64 `json:"soil_moisture_27_to_81cm,omitempty"`
	UvIndex                  []*float64 `json:"uv_index"`
	UvIndexClearSky          []*float64 `json:"uv_index_clear_sky"`
	IsDayOrNight             []*int     `json:"is_day"`
	Cape                     []*float64 `json:"cape"`
	FreezingLevelHeight      []*float64 `json:"freezing_level_height"`
	SunshineDuration         []*float64 `json:"sunshine_duration"`
}

// Forecast retrieve hourly weather based on provide location and requested params
func (service *HourlyWeatherService) Forecast(ctx context.Context, opts *HourlyOptions) (*HourlyWeatherResponse, error) {
	u, err := addOptions("forecast/", opts)

	if err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	result := new(HourlyWeatherResponse)

	_, err = service.client.Do(ctx, req, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
