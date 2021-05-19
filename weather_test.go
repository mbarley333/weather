package weather_test

import (
	"fmt"
	"testing"
	"weather"
)

func TestWeatherGet(t *testing.T) {
	data, err := weather.Get("Kaneohe")
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("No data for you")
	}
	//fmt.Println(data)
}

func TestWeatherDecode(t *testing.T) {
	data, err := weather.Get("Kaneohe")
	if err != nil {
		t.Fatal(err)
	}
	want := "{"coord":{"lon":-157.8036,"lat":21.4181},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":297.42,"feels_like":297.7,"temp_min":295.93,"temp_max":299.15,"pressure":1019,"humidity":69},"visibility":10000,"wind":{"speed":5.66,"deg":80,"gust":8.23},"clouds":{"all":90},"dt":1621451472,"sys":{"type":1,"id":7877,"country":"US","sunrise":1621439481,"sunset":1621487061},"timezone":-36000,"id":5848189,"name":"Kaneohe","cod":200}"
	got := weather.Decode(data)

	fmt.Println(got)

}

func TestWeatherSet(t *testing.T) {
	var want = weather.Weather{
		City:        "Kaneohe",
		Country:     "USA",
		Temperature: 78.0,
	}
	data, err := weather.Get("Kaneohe")
	if err != nil {
		t.Fatal(err)
	}
	got := weather.Set(data)
	if want != got {
		t.Errorf("Weather struct does not match %v %v", want, got)
	}

}
