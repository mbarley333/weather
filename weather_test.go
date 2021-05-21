package weather_test

import (
	"net/http"
	"testing"
	"weather"

	"github.com/jarcoal/httpmock"
)

var response = []byte{123, 34, 99, 111, 111, 114, 100, 34, 58, 123, 34, 108, 111, 110, 34, 58, 45, 49, 53, 55, 46, 56, 48, 51, 54, 44, 34, 108, 97, 116, 34, 58, 50, 49, 46, 52, 49, 56, 49, 125, 44, 34, 119, 101, 97, 116, 104, 101, 114, 34, 58, 91, 123, 34, 105, 100, 34, 58, 56, 48, 51, 44, 34, 109, 97, 105, 110, 34, 58, 34, 67, 108, 111, 117, 100, 115, 34, 44, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 34, 58, 34, 98, 114, 111, 107, 101, 110, 32, 99, 108, 111, 117, 100, 115, 34, 44, 34, 105, 99, 111, 110, 34, 58, 34, 48, 52, 110, 34, 125, 93, 44, 34, 98, 97, 115, 101, 34, 58, 34, 115, 116, 97, 116, 105, 111, 110, 115, 34, 44, 34, 109, 97, 105, 110, 34, 58, 123, 34, 116, 101, 109, 112, 34, 58, 50, 57, 54, 46, 49, 52, 44, 34, 102, 101, 101, 108, 115, 95, 108, 105, 107, 101, 34, 58, 50, 57, 54, 46, 52, 44, 34, 116, 101, 109, 112, 95, 109, 105, 110, 34, 58, 50, 57, 53, 46, 54, 44, 34, 116, 101, 109, 112, 95, 109, 97, 120, 34, 58, 50, 57, 54, 46, 57, 50, 44, 34, 112, 114, 101, 115, 115, 117, 114, 101, 34, 58, 49, 48, 49, 57, 44, 34, 104, 117, 109, 105, 100, 105, 116, 121, 34, 58, 55, 51, 125, 44, 34, 118, 105, 115, 105, 98, 105, 108, 105, 116, 121, 34, 58, 49, 48, 48, 48, 48, 44, 34, 119, 105, 110, 100, 34, 58, 123, 34, 115, 112, 101, 101, 100, 34, 58, 53, 46, 54, 54, 44, 34, 100, 101, 103, 34, 58, 56, 48, 44, 34, 103, 117, 115, 116, 34, 58, 56, 46, 50, 51, 125, 44, 34, 99, 108, 111, 117, 100, 115, 34, 58, 123, 34, 97, 108, 108, 34, 58, 55, 53, 125, 44, 34, 100, 116, 34, 58, 49, 54, 50, 49, 52, 57, 52, 55, 53, 50, 44, 34, 115, 121, 115, 34, 58, 123, 34, 116, 121, 112, 101, 34, 58, 49, 44, 34, 105, 100, 34, 58, 55, 56, 55, 55, 44, 34, 99, 111, 117, 110, 116, 114, 121, 34, 58, 34, 85, 83, 34, 44, 34, 115, 117, 110, 114, 105, 115, 101, 34, 58, 49, 54, 50, 49, 52, 51, 57, 52, 56, 49, 44, 34, 115, 117, 110, 115, 101, 116, 34, 58, 49, 54, 50, 49, 52, 56, 55, 48, 54, 49, 125, 44, 34, 116, 105, 109, 101, 122, 111, 110, 101, 34, 58, 45, 51, 54, 48, 48, 48, 44, 34, 105, 100, 34, 58, 53, 56, 52, 56, 49, 56, 57, 44, 34, 110, 97, 109, 101, 34, 58, 34, 75, 97, 110, 101, 111, 104, 101, 34, 44, 34, 99, 111, 100, 34, 58, 50, 48, 48, 125}

func TestWeatherGet(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.openweathermap.org/data/2.5/weather?q=Kaneohe&appid=3b814c61996538f2e8a2b921e23bbb0a",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, response)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	location := "Kaneohe"
	data, err := weather.Get(location)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("No data for you")
	}

}

func TestWeatherDecode(t *testing.T) {

	want := string(response)
	got, err := weather.Decode(response)
	if err != nil {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

// func TestWeatherSet(t *testing.T) {
// 	var want = weather.Weather{
// 		City:        "Kaneohe",
// 		Country:     "USA",
// 		Temperature: 78.0,
// 	}
// 	data, err := weather.Get("Kaneohe")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	got := weather.Set(data)
// 	if want != got {
// 		t.Errorf("Weather struct does not match %v %v", want, got)
// 	}

// }

// func stash() {
// 	httpmock.RegisterResponder("GET", "https://api.openweathermap.org/data/2.5/weather?q=kaneohe&appid=3b814c61996538f2e8a2b921e23bbb0a",
// 		func(req *http.Request) (*http.Response, error) {
// 			resp := httpmock.NewStringResponse(200,
// 				`{
// 				"coord":{"lon":-157.8036,"lat":21.4181},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"base":"stations","main":{"temp":298.95,"feels_like":299.18,"temp_min":298.9,"temp_max":300.23,"pressure":1017,"humidity":61},"visibility":10000,"wind":{"speed":7.72,"deg":70},"clouds":{"all":20},"dt":1621471005,"sys":{"type":1,"id":7877,"country":"US","sunrise":1621439481,"sunset":1621487061},"timezone":-36000,"id":5848189,"name":"Kaneohe","cod":200
// 			}
// 			`,
// 			)

// 			resp.Header.Add("Content-Type", "application/json")
// 			return resp, nil
// 		},
// 	)
// }
