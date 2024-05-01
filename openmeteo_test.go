package openmeteogo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints. See issue #752.
	baseURLPath = "/test"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the client being tested and is
	// configured to use test server.
	client = NewClient(nil)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.WeatherBaseURL = url

	return client, mux, server.URL, server.Close
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	want, got := defaultWeatherBaseURL, c.WeatherBaseURL.String()
	if want != got {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)

	}

	want, got = defaultUserAgent, c.UserAgent

	if want != got {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestAddOptions(t *testing.T) {
	want := "openmeteo.com?latitude=10.42&longitude=33.52"
	got, _ := addOptions("openmeteo.com", &CurrentOptions{Latitude: 10.42, Longitude: 33.52})

	if want != got {
		t.Errorf("addOptions is different from expected %v, want %v", got, want)
	}

	_, err := addOptions("openmeteocom", "")

	if err == nil {
		t.Errorf("addOptions need to return error")
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	url, _ := url.Parse(defaultWeatherBaseURL)
	inURL, outURL := "/v1/forecast", defaultWeatherBaseURL+"forecast"
	inBody, outBody := &CurrentWeatherResponse{Latitude: 44.32, Longitude: 44.32}, `{"latitude":44.32,"longitude":44.32,"generationtime_ms":0,"utc_offset_seconds":0,"timezone":"","timezone_abbreviation":"","elevation":0,"current_units":null,"current":null}`+"\n"
	req, _ := c.NewRequest("GET", url, inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%v) Body is %v, want %v", inBody, got, want)
	}

	userAgent := req.Header.Get("User-Agent")

	// test that default user-agent is attached to the request
	if got, want := userAgent, c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}

	if !strings.Contains(userAgent, Version) {
		t.Errorf("NewRequest() User-Agent should contain %v, found %v", Version, userAgent)
	}

	inURL = ":openmeteo.com"
	_, err := c.NewRequest("GET", url, inURL, inBody)

	if err == nil {
		t.Errorf("NewRequest() Expected error, found %v", err)
	}
}

func TestBareDo_returnsOpenBody(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	expectedBody := "It's Work!"

	mux.HandleFunc("/test-url", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedBody)
	})

	ctx := context.Background()
	req, err := client.NewRequest("GET", client.WeatherBaseURL, "test-url", nil)
	if err != nil {
		t.Fatalf("client.NewRequest returned error: %v", err)
	}

	resp, err := client.BareDo(ctx, req)
	if err != nil {
		t.Fatalf("client.BareDo returned error: %v", err)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("io.ReadAll returned error: %v", err)
	}
	if string(got) != expectedBody {
		t.Fatalf("Expected %q, got %q", expectedBody, string(got))
	}
	if err := resp.Body.Close(); err != nil {
		t.Fatalf("resp.Body.Close() returned error: %v", err)
	}
}

func TestDo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	type foo struct {
		Foo string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"Foo":"foo"}`)
	})

	req, _ := client.NewRequest("GET", client.WeatherBaseURL, ".", nil)
	body := new(foo)
	ctx := context.Background()
	_, err := client.Do(ctx, req, body)

	if !cmp.Equal(err, nil) {
		t.Errorf("Response err = %v, want %v", body, err)
	}

	want := &foo{"foo"}
	if !cmp.Equal(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", client.WeatherBaseURL, ".", nil)
	ctx := context.Background()
	resp, err := client.Do(ctx, req, nil)

	if err == nil {
		t.Fatal("Expected HTTP 400 error, got no error.")
	}
	if resp.StatusCode != 400 {
		t.Errorf("Expected HTTP 400 error, got %d status code.", resp.StatusCode)
	}
}
