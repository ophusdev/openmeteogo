package openmeteogo

import (
	"context"
)

type HourlyAirQualityService service

const (
	HourlyAirQualityPm10                OpenMeteoConst = "pm10"
	HourlyAirQualityPm25                OpenMeteoConst = "pm2_5"
	HourlyAirQualityCarbonMonoxide      OpenMeteoConst = "carbon_monoxide"
	HourlyAirQualityNitrogenDioxide     OpenMeteoConst = "nitrogen_dioxide"
	HourlyAirQualitySulphurDioxide      OpenMeteoConst = "sulphur_dioxide"
	HourlyAirQualityOzone               OpenMeteoConst = "ozone"
	HourlyAirQualityAerosolOpticalDepth OpenMeteoConst = "aerosol_optical_depth"
	HourlyAirQualityDust                OpenMeteoConst = "dust"
	HourlyAirQualityUvIndex             OpenMeteoConst = "uv_index"
	HourlyAirQualityUvIndexClearSky     OpenMeteoConst = "uv_index_clear_sky"
	HourlyAirQualityAmmonia             OpenMeteoConst = "ammonia"
	HourlyAirQualityAlderPollen         OpenMeteoConst = "alder_pollen"
	HourlyAirQualityBirchPollen         OpenMeteoConst = "birch_pollen"
	HourlyAirQualityGrassPollen         OpenMeteoConst = "grass_pollen"
	HourlyAirQualityMugwortPollen       OpenMeteoConst = "mugwort_pollen"
	HourlyAirQualityOlivePollen         OpenMeteoConst = "olive_pollen"
	HourlyAirQualityRagweedPollen       OpenMeteoConst = "ragweed_pollen"

	HourlyAirQualityEuropeanAqi                OpenMeteoConst = "european_aqi"
	HourlyAirQualityEuropeanAqiPm25            OpenMeteoConst = "european_aqi_pm2_5"
	HourlyAirQualityEuropeanAqiPM10            OpenMeteoConst = "european_aqi_pm10"
	HourlyAirQualityEuropeanAqiNitorgenDioxide OpenMeteoConst = "european_aqi_nitrogen_dioxide"
	HourlyAirQualityEuropeanAqiOzone           OpenMeteoConst = "european_aqi_ozone"
	HourlyAirQualityEuropeanAqiSulphurDioxide  OpenMeteoConst = "european_aqi_sulphur_dioxide"
)

type HourlyAirQualityOptions struct {
	TimeFormat   OpenMeteoConst    `url:"time_format,omitempty"`
	Domains      OpenMeteoConst    `url:"domains,omitempty"`
	Latitude     float64           `url:"latitude"`
	Longitude    float64           `url:"longitude"`
	Hourly       *[]OpenMeteoConst `url:"hourly"`
	ForecastDays int               `url:"forecast_days,omitempty"`
	PastDays     int               `url:"past_days,omitempty"`
}

type HourlyAirQualityResponse struct {
	Latitude             float64                        `json:"latitude"`
	Longitude            float64                        `json:"longitude"`
	Generationtime_ms    float64                        `json:"generationtime_ms"`
	UtcOffsetSeconds     int                            `json:"utc_offset_seconds"`
	Timezone             string                         `json:"timezone"`
	TimezoneAbbreviation string                         `json:"timezone_abbreviation"`
	Elevation            float64                        `json:"elevation"`
	HourlyUnits          HourlyUnitsAirQualityResponse  `json:"hourly_units"`
	Hourly               HourlyRecordAirQualityResponse `json:"hourly"`
}

type HourlyUnitsAirQualityResponse struct {
	Time                       string  `json:"time"`
	Pm10                       *string `json:"pm10,omitempty"`
	Pm25                       *string `json:"pm2_5,omitempty"`
	CarbonMonoxide             *string `json:"carbon_monoxide,omitempty"`
	NitrogenDioxide            *string `json:"nitrogen_dioxide,omitempty"`
	SulphurDioxide             *string `json:"sulphur_dioxide,omitempty"`
	Ozone                      *string `json:"ozone,omitempty"`
	AerosolOpticalDepth        *string `json:"aerosol_optical_depth,omitempty"`
	Dust                       *string `json:"dust,omitempty"`
	UvIndex                    *string `json:"uv_index,omitempty"`
	UvIndexClearSky            *string `json:"uv_index_clear_sky,omitempty"`
	Ammonia                    *string `json:"ammonia,omitempty"`
	AlderPollen                *string `json:"alder_pollen,omitempty"`
	BirchPollen                *string `json:"birch_pollen,omitempty"`
	GrassPollen                *string `json:"grass_pollen,omitempty"`
	MugwortPollen              *string `json:"mugwort_pollen,omitempty"`
	OlivePollen                *string `json:"olive_pollen,omitempty"`
	RagweedPollen              *string `json:"ragweed_pollen,omitempty"`
	EuropeanAqi                *string `json:"european_aqi,omitempty"`
	EuropeanAqiPm25            *string `json:"european_aqi_pm2_5,omitempty"`
	EuropeanAqiPM10            *string `json:"european_aqi_pm10,omitempty"`
	EuropeanAqiNitorgenDioxide *string `json:"european_aqi_nitrogen_dioxide,omitempty"`
	EuropeanAqiOzone           *string `json:"european_aqi_ozone,omitempty"`
	EuropeanAqiSulphurDioxide  *string `json:"european_aqi_sulphur_dioxide,omitempty"`
}

type HourlyRecordAirQualityResponse struct {
	Time                       []string   `json:"time,omitempty"`
	Pm10                       []*float64 `json:"pm10,omitempty"`
	Pm25                       []*float64 `json:"pm2_5,omitempty"`
	CarbonMonoxide             []*float64 `json:"carbon_monoxide,omitempty"`
	NitrogenDioxide            []*float64 `json:"nitrogen_dioxide,omitempty"`
	SulphurDioxide             []*float64 `json:"sulphur_dioxide,omitempty"`
	Ozone                      []*float64 `json:"ozone,omitempty"`
	AerosolOpticalDepth        []*float64 `json:"aerosol_optical_depth,omitempty"`
	Dust                       []*float64 `json:"dust,omitempty"`
	UvIndex                    []*float64 `json:"uv_index,omitempty"`
	UvIndexClearSky            []*float64 `json:"uv_index_clear_sky,omitempty"`
	Ammonia                    []*float64 `json:"ammonia,omitempty"`
	AlderPollen                []*float64 `json:"alder_pollen,omitempty"`
	BirchPollen                []*float64 `json:"birch_pollen,omitempty"`
	GrassPollen                []*float64 `json:"grass_pollen,omitempty"`
	MugwortPollen              []*float64 `json:"mugwort_pollen,omitempty"`
	OlivePollen                []*float64 `json:"olive_pollen,omitempty"`
	RagweedPollen              []*float64 `json:"ragweed_pollen,omitempty"`
	EuropeanAqi                []*float64 `json:"european_aqi,omitempty"`
	EuropeanAqiPm25            []*float64 `json:"european_aqi_pm2_5,omitempty"`
	EuropeanAqiPM10            []*float64 `json:"european_aqi_pm10,omitempty"`
	EuropeanAqiNitorgenDioxide []*float64 `json:"european_aqi_nitrogen_dioxide,omitempty"`
	EuropeanAqiOzone           []*float64 `json:"european_aqi_ozone,omitempty"`
	EuropeanAqiSulphurDioxide  []*float64 `json:"european_aqi_sulphur_dioxide,omitempty"`
}

// Forecast retrieve hourly air quality based on provide location and requested params
func (service *HourlyAirQualityService) Forecast(ctx context.Context, opts *HourlyAirQualityOptions) (*HourlyAirQualityResponse, error) {
	u, err := addOptions("air-quality", opts)

	if err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("GET", service.client.AirQualityBaseURL, u, nil)

	if err != nil {
		return nil, err
	}

	result := new(HourlyAirQualityResponse)

	_, err = service.client.Do(ctx, req, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
