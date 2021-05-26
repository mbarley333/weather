package weather_test

import (
	"fmt"
	"net/http"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jarcoal/httpmock"
)

var response = []byte{123, 34, 99, 111, 111, 114, 100, 34, 58, 123, 34, 108, 111, 110, 34, 58, 45, 49, 53, 55, 46, 56, 48, 51, 54, 44, 34, 108, 97, 116, 34, 58, 50, 49, 46, 52, 49, 56, 49, 125, 44, 34, 119, 101, 97, 116, 104, 101, 114, 34, 58, 91, 123, 34, 105, 100, 34, 58, 56, 48, 51, 44, 34, 109, 97, 105, 110, 34, 58, 34, 67, 108, 111, 117, 100, 115, 34, 44, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 34, 58, 34, 98, 114, 111, 107, 101, 110, 32, 99, 108, 111, 117, 100, 115, 34, 44, 34, 105, 99, 111, 110, 34, 58, 34, 48, 52, 110, 34, 125, 93, 44, 34, 98, 97, 115, 101, 34, 58, 34, 115, 116, 97, 116, 105, 111, 110, 115, 34, 44, 34, 109, 97, 105, 110, 34, 58, 123, 34, 116, 101, 109, 112, 34, 58, 50, 57, 54, 46, 49, 52, 44, 34, 102, 101, 101, 108, 115, 95, 108, 105, 107, 101, 34, 58, 50, 57, 54, 46, 52, 44, 34, 116, 101, 109, 112, 95, 109, 105, 110, 34, 58, 50, 57, 53, 46, 54, 44, 34, 116, 101, 109, 112, 95, 109, 97, 120, 34, 58, 50, 57, 54, 46, 57, 50, 44, 34, 112, 114, 101, 115, 115, 117, 114, 101, 34, 58, 49, 48, 49, 57, 44, 34, 104, 117, 109, 105, 100, 105, 116, 121, 34, 58, 55, 51, 125, 44, 34, 118, 105, 115, 105, 98, 105, 108, 105, 116, 121, 34, 58, 49, 48, 48, 48, 48, 44, 34, 119, 105, 110, 100, 34, 58, 123, 34, 115, 112, 101, 101, 100, 34, 58, 53, 46, 54, 54, 44, 34, 100, 101, 103, 34, 58, 56, 48, 44, 34, 103, 117, 115, 116, 34, 58, 56, 46, 50, 51, 125, 44, 34, 99, 108, 111, 117, 100, 115, 34, 58, 123, 34, 97, 108, 108, 34, 58, 55, 53, 125, 44, 34, 100, 116, 34, 58, 49, 54, 50, 49, 52, 57, 52, 55, 53, 50, 44, 34, 115, 121, 115, 34, 58, 123, 34, 116, 121, 112, 101, 34, 58, 49, 44, 34, 105, 100, 34, 58, 55, 56, 55, 55, 44, 34, 99, 111, 117, 110, 116, 114, 121, 34, 58, 34, 85, 83, 34, 44, 34, 115, 117, 110, 114, 105, 115, 101, 34, 58, 49, 54, 50, 49, 52, 51, 57, 52, 56, 49, 44, 34, 115, 117, 110, 115, 101, 116, 34, 58, 49, 54, 50, 49, 52, 56, 55, 48, 54, 49, 125, 44, 34, 116, 105, 109, 101, 122, 111, 110, 101, 34, 58, 45, 51, 54, 48, 48, 48, 44, 34, 105, 100, 34, 58, 53, 56, 52, 56, 49, 56, 57, 44, 34, 110, 97, 109, 101, 34, 58, 34, 75, 97, 110, 101, 111, 104, 101, 34, 44, 34, 99, 111, 100, 34, 58, 50, 48, 48, 125}

func TestWeatherGet(t *testing.T) {

	api, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		t.Fatal(err)
	}
	location := "Kaneohe"

	url, err := weather.SetApiURL(location, api)
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Weather{
		Main:        "Clouds",
		Description: "broken clouds",
		TempK:       296.14,
		City:        "Kaneohe",
		Latitude:    21.4181,
		Longitude:   -157.8036,
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url, httpmock.NewBytesResponder(http.StatusOK, response))

	got, err := weather.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(weather.WeatherResponse{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestWeatherSetApiURL(t *testing.T) {
	api, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		t.Fatal(err)
	}

	location := "Kaneohe"

	url, err := weather.SetApiURL(location, api)
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, api)
	got := url
	if want != got {
		t.Errorf("\r\nwanted: %s\r\n, \r\ngot: %s", want, got)
	}
}

func TestGetWeatherAPIKey(t *testing.T) {

	_, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		t.Fatal(err)
	}

}

func TestConvertTempF(t *testing.T) {

	w := weather.Weather{

		TempK: 296.14,
	}

	want := 73.38200000000002

	w.SetTemp(w.TempK)
	got := w.TempF

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(weather.WeatherResponse{})) {
		t.Error(cmp.Diff(want, got))
	}

	want = 22.99000000000001
	got = w.TempC

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(weather.WeatherResponse{})) {
		t.Error(cmp.Diff(want, got))
	}

}
