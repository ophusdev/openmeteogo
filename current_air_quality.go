package openmeteogo

import (
	"context"
)

type CurrentAirQualityService service

const (
	CurrentAirQualityPm10                OpenMeteoConst = "pm10"
	CurrentAirQualityPm25                OpenMeteoConst = "pm2_5"
	CurrentAirQualityCarbonMonoxide      OpenMeteoConst = "carbon_monoxide"
	CurrentAirQualityNitrogenDioxide     OpenMeteoConst = "nitrogen_dioxide"
	CurrentAirQualitySulphurDioxide      OpenMeteoConst = "sulphur_dioxide"
	CurrentAirQualityOzone               OpenMeteoConst = "ozone"
	CurrentAirQualityAerosolOpticalDepth OpenMeteoConst = "aerosol_optical_depth"
	CurrentAirQualityDust                OpenMeteoConst = "dust"
	CurrentAirQualityUvIndex             OpenMeteoConst = "uv_index"
	CurrentAirQualityUvIndexClearSky     OpenMeteoConst = "uv_index_clear_sky"
	CurrentAirQualityAmmonia             OpenMeteoConst = "ammonia"
	CurrentAirQualityAlderPollen         OpenMeteoConst = "alder_pollen"
	CurrentAirQualityBirchPollen         OpenMeteoConst = "birch_pollen"
	CurrentAirQualityGrassPollen         OpenMeteoConst = "grass_pollen"
	CurrentAirQualityMugwortPollen       OpenMeteoConst = "mugwort_pollen"
	CurrentAirQualityOlivePollen         OpenMeteoConst = "olive_pollen"
	CurrentAirQualityRagweedPollen       OpenMeteoConst = "ragweed_pollen"
	CurrentAirQualityEuropeanAqi         OpenMeteoConst = "european_aqi"
	CurrentAirQualityUsAqi               OpenMeteoConst = "us_aqi"
)

type CurrentAirQualityOptions struct {
	TimeFormat   OpenMeteoConst    `url:"time_format,omitempty"`
	Domains      OpenMeteoConst    `url:"domains,omitempty"`
	Latitude     float64           `url:"latitude"`
	Longitude    float64           `url:"longitude"`
	Current      *[]OpenMeteoConst `url:"current"`
	ForecastDays int               `url:"forecast_days,omitempty"`
	PastDays     int               `url:"past_days,omitempty"`
}

type CurrentAirQualityResponse struct {
	Latitude             float64                         `json:"latitude"`
	Longitude            float64                         `json:"longitude"`
	Generationtime_ms    float64                         `json:"generationtime_ms"`
	UtcOffsetSeconds     int                             `json:"utc_offset_seconds"`
	Timezone             string                          `json:"timezone"`
	TimezoneAbbreviation string                          `json:"timezone_abbreviation"`
	Elevation            float64                         `json:"elevation"`
	CurrentUnits         CurrentUnitsAirQualityResponse  `json:"current_units"`
	Current              CurrentRecordAirQualityResponse `json:"current"`
}

type CurrentUnitsAirQualityResponse struct {
	Time                string  `json:"time"`
	Interval            string  `json:"interval"`
	Pm10                *string `json:"pm10,omitempty"`
	Pm25                *string `json:"pm2_5,omitempty"`
	CarbonMonoxide      *string `json:"carbon_monoxide,omitempty"`
	NitrogenDioxide     *string `json:"nitrogen_dioxide,omitempty"`
	SulphurDioxide      *string `json:"sulphur_dioxide,omitempty"`
	Ozone               *string `json:"ozone,omitempty"`
	AerosolOpticalDepth *string `json:"aerosol_optical_depth,omitempty"`
	Dust                *string `json:"dust,omitempty"`
	UvIndex             *string `json:"uv_index,omitempty"`
	UvIndexClearSky     *string `json:"uv_index_clear_sky,omitempty"`
	Ammonia             *string `json:"ammonia,omitempty"`
	AlderPollen         *string `json:"alder_pollen,omitempty"`
	BirchPollen         *string `json:"birch_pollen,omitempty"`
	GrassPollen         *string `json:"grass_pollen,omitempty"`
	MugwortPollen       *string `json:"mugwort_pollen,omitempty"`
	OlivePollen         *string `json:"olive_pollen,omitempty"`
	RagweedPollen       *string `json:"ragweed_pollen,omitempty"`
	EuropeanAqi         *string `json:"european_aqi,omitempty"`
	UsAqi               *string `json:"us_aqi,omitempty"`
}

type CurrentRecordAirQualityResponse struct {
	Time                string   `json:"time,omitempty"`
	Interval            int      `json:"interval,omitempty"`
	Pm10                *float64 `json:"pm10,omitempty"`
	Pm25                *float64 `json:"pm2_5,omitempty"`
	CarbonMonoxide      *float64 `json:"carbon_monoxide,omitempty"`
	NitrogenDioxide     *float64 `json:"nitrogen_dioxide,omitempty"`
	SulphurDioxide      *float64 `json:"sulphur_dioxide,omitempty"`
	Ozone               *float64 `json:"ozone,omitempty"`
	AerosolOpticalDepth *float64 `json:"aerosol_optical_depth,omitempty"`
	Dust                *float64 `json:"dust,omitempty"`
	UvIndex             *float64 `json:"uv_index,omitempty"`
	UvIndexClearSky     *float64 `json:"uv_index_clear_sky,omitempty"`
	Ammonia             *float64 `json:"ammonia,omitempty"`
	AlderPollen         *float64 `json:"alder_pollen,omitempty"`
	BirchPollen         *float64 `json:"birch_pollen,omitempty"`
	GrassPollen         *float64 `json:"grass_pollen,omitempty"`
	MugwortPollen       *float64 `json:"mugwort_pollen,omitempty"`
	OlivePollen         *float64 `json:"olive_pollen,omitempty"`
	RagweedPollen       *float64 `json:"ragweed_pollen,omitempty"`
	EuropeanAqi         *float64 `json:"european_aqi,omitempty"`
	UsAqi               *float64 `json:"us_aqi,omitempty"`
}

// Forecast retrieve Current air quality based on provide location and requested params
func (service *CurrentAirQualityService) Forecast(ctx context.Context, opts *CurrentAirQualityOptions) (*CurrentAirQualityResponse, error) {
	u, err := addOptions("air-quality", opts)

	if err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("GET", service.client.AirQualityBaseURL, u, nil)

	if err != nil {
		return nil, err
	}

	result := new(CurrentAirQualityResponse)

	_, err = service.client.Do(ctx, req, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
