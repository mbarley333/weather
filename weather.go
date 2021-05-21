package weather

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	City        string  `json:"name"`
	WeatherMain string  `json:"main"`
	Humidity    string  `json:"humidity"`
	Temperature float64 `json:"temp"`
}

//{"coord":{"lon":-157.8036,"lat":21.4181},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"base":"stations","main":{"temp":296.14,"feels_like":296.4,"temp_min":295.6,"temp_max":296.92,"pressure":1019,"humidity":73},"visibility":10000,"wind":{"speed":5.66,"deg":80,"gust":8.23},"clouds":{"all":75},"dt":1621494752,"sys":{"type":1,"id":7877,"country":"US","sunrise":1621439481,"sunset":1621487061},"timezone":-36000,"id":5848189,"name":"Kaneohe","cod":200}PASS
func Get(location string) ([]byte, error) {

	//get api response
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=3b814c61996538f2e8a2b921e23bbb0a", location))

	if err != nil {
		return nil, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	//fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code.  %v", resp.StatusCode)
	}

	//fmt.Printf("%+v", resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	//fmt.Printf("%s", data)

	return data, nil
}

func Decode(data []byte) (string, error) {

	if len(data) == 0 {
		return "", fmt.Errorf("empty byte slice")
	}
	return string(data), nil
}

func Set(weather []byte) Weather {

	return Weather{}
}
