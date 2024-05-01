package openmeteogo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestHourlyAirQualityResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	remoteResponse := `{
		"latitude": 52.549995,
		"longitude": 13.450001,
		"generationtime_ms": 0.04601478576660156,
		"utc_offset_seconds": 0,
		"timezone": "GMT",
		"timezone_abbreviation": "GMT",
		"elevation": 38,
		"hourly_units": {"time": "iso8601","pm10": "μg/m³","pm2_5": "μg/m³"},
		"hourly": {"time": ["2024-05-01T00:00","2024-05-01T01:00"],"pm10": [8.4,8.6],"pm2_5": [5.7,6]}
	}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, remoteResponse)
	})

	ctx := context.Background()

	opts := &HourlyAirQualityOptions{
		Latitude:     52.52,
		Longitude:    13.41,
		ForecastDays: 1,
		Hourly: &[]OpenMeteoConst{
			HourlyAirQualityPm10,
			HourlyAirQualityPm25,
		},
	}

	gotResponse, err := client.HourlyAirQuality.Forecast(ctx, opts)

	if err != nil {
		t.Fatal("Unexpected error")
	}

	wantResponse := &HourlyAirQualityResponse{}

	_ = json.Unmarshal([]byte(remoteResponse), wantResponse)

	if wantResponse.Latitude != gotResponse.Latitude {
		t.Errorf("Response body = %v, want %v", gotResponse, wantResponse)
	}
}
