package weather

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	City        string
	Country     string
	Temperature float64
}

func Get(location string) ([]byte, error) {

	//get api response
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=3b814c61996538f2e8a2b921e23bbb0a", location))
	if err != nil {
		return nil, fmt.Errorf("Something went wrong.  Please try again later.  %v", err)
	}

	//fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code.  %v", resp.StatusCode)
	}

	//fmt.Printf("%+v", resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong.  Please try again later.  %v", err)
	}
	fmt.Printf("%s", data)

	return data, nil
}

func Decode(weather []byte) string {

	return ""
}

func Set(weather []byte) Weather {

	return Weather{}
}
