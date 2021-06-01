package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestWeatherGet(t *testing.T) {
	t.Parallel()
	apiKey := "dummy"
	tempUnits := "imperial"
	location := "Kaneohe"

	//setup http server for get requests
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/weather_test.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		w.WriteHeader(http.StatusOK)
		io.Copy(w, f)

	}))

	//create new client based on struct
	client, err := weather.NewClient(apiKey, tempUnits)
	if err != nil {
		t.Fatal(err)
	}

	//set base url to test server url
	client.Base = ts.URL

	//set HTTPClient to test client to handle x509 certs w/o more setup work
	client.HTTPClient = ts.Client()
	got, err := client.Get(location)
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Weather{
		Main:        "Clouds",
		Description: "broken clouds",
		Temp:        74.12,
		City:        "Kaneohe",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetWeatherAPIKey(t *testing.T) {
	t.Parallel()
	_, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		t.Fatal(err)
	}

}

func TestNewClient(t *testing.T) {
	t.Parallel()
	apiKey := "dummy"

	tempUnits := "imperial"

	got, err := weather.NewClient(apiKey, tempUnits)
	if err != nil {
		t.Fatal(err)
	}

	//want api key and temp units
	want := weather.Client{
		Base:       "https://api.openweathermap.org",
		ApiKey:     apiKey,
		Units:      tempUnits,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestFormatURL(t *testing.T) {
	t.Parallel()

	apiKey := "dummy"
	tempUnits := "imperial"
	location := "Kaneohe"
	client, err := weather.NewClient(apiKey, tempUnits)
	if err != nil {
		t.Fatal(err)
	}

	want := "https://api.openweathermap.org/data/2.5/weather?q=Kaneohe&units=imperial&appid=dummy"
	got := client.FormatURL(location)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestUnmarshallJson(t *testing.T) {
	t.Parallel()
	f, err := os.Open("testdata/weather_test.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	got, err := weather.ParseResponse(f)
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Weather{
		Main:        "Clouds",
		Description: "broken clouds",
		Temp:        74.12,
		City:        "Kaneohe",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
