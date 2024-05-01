package openmeteogo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCurrentAirQualityResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	remoteResponse := `{
        "latitude": 52.549995,
        "longitude": 13.450001,
        "generationtime_ms": 0.07700920104980469,
        "utc_offset_seconds": 0,
        "timezone": "GMT",
        "timezone_abbreviation": "GMT",
        "elevation": 38,
        "current_units": {
                "time": "iso8601",
                "interval": "seconds",
                "pm10": "μg/m³",
                "pm2_5": "μg/m³"
        },
        "current": {
                "time": "2024-05-01T12:00",
                "interval": 3600,
                "pm10": 7.1,
                "pm2_5": 5.5
        }
	}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, remoteResponse)
	})

	ctx := context.Background()

	opts := &CurrentAirQualityOptions{
		Latitude:     52.52,
		Longitude:    13.41,
		ForecastDays: 1,
		Domains:      AirQualityDomainEurope,
		Current: &[]OpenMeteoConst{
			CurrentAirQualityPm10,
			CurrentAirQualityPm25,
		},
	}

	gotResponse, err := client.CurrentAirQuality.Forecast(ctx, opts)

	if err != nil {
		t.Fatal("Unexpected error")
	}

	wantResponse := &CurrentAirQualityResponse{}

	_ = json.Unmarshal([]byte(remoteResponse), wantResponse)

	if wantResponse.Latitude != gotResponse.Latitude {
		t.Errorf("Response body = %v, want %v", gotResponse, wantResponse)
	}
}
