package openmeteogo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestDailyWeatherResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	remoteResponse := `{"latitude":52.52}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, remoteResponse)
	})

	ctx := context.Background()

	opts := DailyOptions{
		Latitude: 52.52,
	}

	gotResponse, err := client.DailyWeather.Forecast(ctx, &opts)

	if err != nil {
		t.Fatal("Unexpected error")
	}

	wantResponse := DailyWeatherResponse{}
	_ = json.Unmarshal([]byte(remoteResponse), &wantResponse)

	if wantResponse.Latitude != gotResponse.Latitude {
		t.Errorf("Response body = %v, want %v", gotResponse, wantResponse)
	}
}

func TestDailyWeatherResponseError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	ctx := context.Background()

	opts := DailyOptions{
		Latitude: 52.52,
	}

	_, err := client.DailyWeather.Forecast(ctx, &opts)

	if err == nil {
		t.Fatal("Unexpected error")
	}
}
