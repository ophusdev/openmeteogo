package openmeteogo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	Version = "0.3.0"

	defaultBaseURL   = "https://api.open-meteo.com/v1/"
	defaultUserAgent = "openmeteo-go" + "/" + Version
	defaultForecast  = "forecast"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL

	UserAgent string

	CurrentWeather *CurrentWeatherService
	DailyWeather   *DailyWeatherService
	HourlyWeather  *HourlyWeatherService

	// TODO: add Air pollution and other services...

	common service
}

type service struct {
	client *Client
}

type Response struct {
	*http.Response
}

type AcceptedError struct {
	// Raw contains the response body.
	Raw []byte
}

func (*AcceptedError) Error() string {
	return "api error. Try again or check params"
}

// Is returns whether the provided error equals this error.
func (ae *AcceptedError) Is(target error) bool {
	v, ok := target.(*AcceptedError)
	if !ok {
		return false
	}
	return bytes.Equal(ae.Raw, v.Raw)
}

type ErrorResponse struct {
	Response *http.Response `json:"-"`      // HTTP response that caused this error
	Message  string         `json:"reason"` // error message
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

type RequestOption func(req *http.Request)

// NewClient return a new OpenMeteoClient
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient2 := *httpClient
	c := &Client{client: &httpClient2}
	c.initialize()
	return c
}

// initialize Client
func (c *Client) initialize() {
	if c.client == nil {
		c.client = &http.Client{}
	}

	if c.BaseURL == nil {
		c.BaseURL, _ = url.Parse(defaultBaseURL)
	}

	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}

	c.common.client = c

	c.CurrentWeather = (*CurrentWeatherService)(&c.common)
	c.HourlyWeather = (*HourlyWeatherService)(&c.common)
	c.DailyWeather = (*DailyWeatherService)(&c.common)
}

func (client *Client) NewRequest(method string, url string, body interface{}, opts ...RequestOption) (*http.Request, error) {
	u, err := client.BaseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if client.UserAgent != "" {
		req.Header.Set("User-Agent", client.UserAgent)
	}

	return req, nil
}

func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	//TODO: validate ForecastDays is > 0

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func (client *Client) BareDo(context context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(context)

	resp, err := client.client.Do(req)

	if err != nil {
		return nil, err
	}

	response := newResponse(resp)

	err = CheckResponse(resp)

	return response, err
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusAccepted {
		return &AcceptedError{}
	}

	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}

	data, err := io.ReadAll(r.Body)

	if err == nil && data != nil {
		err = json.Unmarshal(data, errorResponse)
		if err != nil {
			// reset the response as if this never happened
			errorResponse = &ErrorResponse{Response: r}
		}
	}

	return errorResponse
}
