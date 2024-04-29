# Openmeteogo

Openmeteogo is a Go library to retrieve meteo data from [open-meteo.com](https://open-meteo.com).

## Installation

```bash
go get github.com/ophusdev/openmeteogo
```

## Usage
Check https://open-meteo.com/en/docs/ for docs and useful params

### Current Weather
```go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ophusdev/openmeteogo"
)

func main() {

	ctx := context.Background()

	client := openmeteogo.NewClient(nil)

	opts := &openmeteogo.CurrentOptions{
		Latitude:  41.902782,
		Longitude: 12.496366,
		Current: &[]openmeteogo.OpenMeteoConst{
			openmeteogo.CurrentTemperature2m,
			openmeteogo.CurrentRelativeHumidity2m,
			openmeteogo.CurrentApparentTemperature,
			openmeteogo.CurrentIsDay,
			openmeteogo.CurrentPrecipitation,
			openmeteogo.CurrentRain,
			openmeteogo.CurrentShowers,
			openmeteogo.CurrentSnowfall,
			openmeteogo.CurrentCloudCover,
			openmeteogo.CurrentWeatherCode,
			openmeteogo.CurrentSealevelPressure,
			openmeteogo.CurrentSurfacePressure,
			openmeteogo.CurrentWindSpeed10m,
			openmeteogo.CurrentWindDirection10m,
			openmeteogo.CurrentWindGust10m,
		},
	}

	forecast, err := client.CurrentWeather.Forecast(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	s, _ := json.MarshalIndent(forecast, "", "\t")

	fmt.Print(string(s))

}

```

### Hourly Forecast
```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ophusdev/openmeteogo"
)

func main() {

	ctx := context.Background()

	client := openmeteogo.NewClient(nil)

	opts := &openmeteogo.HourlyOptions{
		Latitude:     41.902782,
		Longitude:    12.496366,
		ForecastDays: 3,
		Hourly: &[]openmeteogo.OpenMeteoConst{
			openmeteogo.HourlyTemperature2m,
			openmeteogo.HourlyRelativeHumidity2m,
			openmeteogo.HourlyDewpoint2m,
			openmeteogo.HourlyApparentTemperature,
			openmeteogo.HourlyPrecipitationProbability,
			openmeteogo.HourlyPrecipitation,
			openmeteogo.HourlyRain,
			openmeteogo.HourlyShowers,
			openmeteogo.HourlySnowfall,
			openmeteogo.HourlySnowDepth,
			openmeteogo.HourlyWeathercode,
			openmeteogo.HourlySealevelPressure,
			openmeteogo.HourlySurfacePressure,
			openmeteogo.HourlyCloudcoverTotal,
			openmeteogo.HourlyCloudcoverLow,
			openmeteogo.HourlyCloudcoverMid,
			openmeteogo.HourlyCloudcoverHigh,
			openmeteogo.HourlyVisibility,
			openmeteogo.HourlyEvapotranspiration,
			openmeteogo.HourlyEt0FaoEvapotranspiration,
			openmeteogo.HourlyVapourPressureDeficit,
			openmeteogo.HourlyWindSpeed10m,
			openmeteogo.HourlyWindSpeed80m,
			openmeteogo.HourlyWindSpeed120m,
			openmeteogo.HourlyWindSpeed180m,
			openmeteogo.HourlyWindDirection10m,
			openmeteogo.HourlyWindDirection80m,
			openmeteogo.HourlyWindDirection120m,
			openmeteogo.HourlyWindDirection180m,
			openmeteogo.HourlyWindGusts10m,
			openmeteogo.HourlyTemperature80m,
			openmeteogo.HourlyTemperature120m,
			openmeteogo.HourlyTemperature180m,
			openmeteogo.HourlySoilTemperature0cm,
			openmeteogo.HourlySoilTemperature6cm,
			openmeteogo.HourlySoilTemperature18cm,
			openmeteogo.HourlySoilTemperature54cm,
			openmeteogo.HourlySoilMoisture0to1cm,
			openmeteogo.HourlySoilMoisture1to3cm,
			openmeteogo.HourlySoilMoisture3to9cm,
			openmeteogo.HourlySoilMoisture9to27cm,
			openmeteogo.HourlySoilMoisture27to81cm,
		},
	}

	forecast, err := client.HourlyWeather.Forecast(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	s, _ := json.MarshalIndent(forecast, "", "\t")

	fmt.Print(string(s))

}

```


### Daily Forecast
```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ophusdev/openmeteogo"
)

func main() {

	ctx := context.Background()

	client := openmeteogo.NewClient(nil)

	opts := &openmeteogo.DailyOptions{
		Latitude:     41.902782,
		Longitude:    12.496366,
		ForecastDays: 7,
		Daily: &[]openmeteogo.OpenMeteoConst{
			openmeteogo.DailyWeatherCode,
			openmeteogo.DailyTemperature2mMax,
			openmeteogo.DailyTemperature2mMin,
			openmeteogo.DailyApparentTemperatureMax,
			openmeteogo.DailyApparentTemperatureMin,
			openmeteogo.DailySunrise,
			openmeteogo.DailySunset,
			openmeteogo.DailyDaylightDuration,
			openmeteogo.DailySunshineDuration,
			openmeteogo.DailyUvIndexMax,
			openmeteogo.DailyUvIndexClearSkyMax,
			openmeteogo.DailyPrecipitationSum,
			openmeteogo.DailyRainSum,
			openmeteogo.DailyShowersSum,
			openmeteogo.DailySnowfallSum,
			openmeteogo.DailyPrecipitationHours,
			openmeteogo.DailyPrecipitationProbabilityMax,
			openmeteogo.DailyWindSpeed10mMax,
			openmeteogo.DailyWindGusts10mMax,
			openmeteogo.DailyWindDirection10mDominant,
			openmeteogo.DailyShortwaveRadiationSum,
			openmeteogo.DailyEt0FaoEvapotranspiration,
		},
	}

	forecast, err := client.DailyWeather.Forecast(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	s, _ := json.MarshalIndent(forecast, "", "\t")

	fmt.Print(string(s))

}

```


## License

[MIT](https://choosealicense.com/licenses/mit/)