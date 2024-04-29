package openmeteogo

type OpenMeteoConst string

const (
	TimezoneGMT        OpenMeteoConst = "GMT"
	TimezoneAnchorage  OpenMeteoConst = "America/Anchorage"
	TimezoneLosAngeles OpenMeteoConst = "America/Los_Angeles"
	TimezoneDenver     OpenMeteoConst = "America/Denver"
	TimezoneChicago    OpenMeteoConst = "America/Chicago"
	TimezoneNewYork    OpenMeteoConst = "America/NewYork"
	TimezoneSaoPaulo   OpenMeteoConst = "America/SaoPaulo"
	TimezoneLondon     OpenMeteoConst = "Europe/London"
	TimezoneBerlin     OpenMeteoConst = "Europe/Berlin"
	TimezoneMoscow     OpenMeteoConst = "Europe/Moscow"
	TimezoneCairo      OpenMeteoConst = "Africa/Cairo"
	TimezoneBangkok    OpenMeteoConst = "Asia/Bangkok"
	TimezoneSingapore  OpenMeteoConst = "Asia/Singapore"
	TimezoneTokio      OpenMeteoConst = "Asia/Tokio"
	TimezoneSydney     OpenMeteoConst = "Australia/Sydney"
	TimezoneAuckland   OpenMeteoConst = "Pacific/Auckland"

	TemperatureUnitCelsius    OpenMeteoConst = "celsius"
	TemperatureUnitFahrenheit OpenMeteoConst = "fahrenheit"

	WindSpeedUnitKmH   OpenMeteoConst = "kmh"
	WindSpeedUnitMs    OpenMeteoConst = "ms"
	WindSpeedUnitMph   OpenMeteoConst = "mph"
	WindSpeedUnitKnots OpenMeteoConst = "kn"

	PrecipitationUnitMm   OpenMeteoConst = "mm"
	PrecipitationUnitInch OpenMeteoConst = "inch"

	TimeFormatIso8601  OpenMeteoConst = "iso8601"
	TimeFormatUnixTime OpenMeteoConst = "unixtime"
)

type Options struct {
	Timezone          OpenMeteoConst `url:"timezone,omitempty"`
	TemperatureUnit   OpenMeteoConst `url:"temperature_unit,omitempty"`
	WindSpeedUnit     OpenMeteoConst `url:"wind_speed_unit,omitempty"`
	PrecipitationUnit OpenMeteoConst `url:"precipitation_unit,omitempty"`
	TimeFormat        OpenMeteoConst `url:"time_format,omitempty"`
}
