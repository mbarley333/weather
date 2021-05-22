package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type weatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	City  string `json:"name"`
	Coord struct {
		Longitude float64 `json:"lon"`
		Latitude  float64 `json:"lat"`
	} `json:"coord"`
}

func Get(location string) (weatherResponse, error) {

	//get api response
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=3b814c61996538f2e8a2b921e23bbb0a", location))

	if err != nil {
		return weatherResponse{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return weatherResponse{}, fmt.Errorf("unexpected status code.  %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherResponse{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	var wdata weatherResponse
	err = json.Unmarshal(data, &wdata)
	if err != nil {
		log.Fatal("unable to unmarshall data")
	}

	return wdata, nil
}
