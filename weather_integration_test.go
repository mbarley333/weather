//+build integration

package weather_test

import (
	"testing"
	"weather"
)

// used to test against actual open weather map api
// initiated by go test -tags=integration
func TestWeatherGetIntegration(t *testing.T) {
	apiKey, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		t.Fatal(err)
	}

	tempUnits := "imperial"
	location := "Kaneohe"
	client, err := weather.NewClient(apiKey, tempUnits)
	if err != nil {
		t.Fatal(err)
	}

	weather, err := client.Get(location)
	if err != nil {
		t.Fatal(err)
	}

	if weather.Main == "" {
		t.Fatalf("Invalid weather struct %v", weather)
	}

}
