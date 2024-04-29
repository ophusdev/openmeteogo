package openmeteogo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCurrentWeatherResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	remoteResponse := `{"latitude":52.52}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, remoteResponse)
	})

	ctx := context.Background()

	opts := CurrentOptions{
		Latitude: 52.52,
	}

	gotResponse, err := client.CurrentWeather.Forecast(ctx, &opts)

	if err != nil {
		t.Fatal("Unexpected error")
	}

	wantResponse := CurrentWeatherResponse{}
	_ = json.Unmarshal([]byte(remoteResponse), &wantResponse)

	if wantResponse != *gotResponse {
		t.Errorf("Response body = %v, want %v", gotResponse, wantResponse)
	}
}

func TestCurrentWeatherResponseError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	ctx := context.Background()

	opts := CurrentOptions{
		Latitude: 52.52,
	}

	_, err := client.CurrentWeather.Forecast(ctx, &opts)

	if err == nil {
		t.Fatal("Unexpected error")
	}
}
