package weather_test

import (
	"testing"
	"weather"
)

func TestWeatherGet(t *testing.T) {
	data, err := weather.Get("London")
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("No data for you")
	}
}
