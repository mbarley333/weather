package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

type ApiURL struct {
	Base     string
	Location string
	ApiKey   string
}

func Get(location string) (weatherResponse, error) {

	api, err := GetWeatherAPIKey("../secret/weather.txt")
	if err != nil {
		return weatherResponse{}, fmt.Errorf("something went wrong getting api key.  Please try again later.  %v", err)
	}

	//get api response
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, api))

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

func SetApiURL(location string) (string, error) {

	if len(location) == 0 {
		return "", fmt.Errorf("something went wrong.  Please provide location")
	}

	api, err := GetWeatherAPIKey("../secret/weather.txt")
	if err != nil {
		log.Fatal(err)
	}

	var aData ApiURL

	aData.Base = "https://api.openweathermap.org/data/2.5/weather?q="
	aData.Location = location
	aData.ApiKey = fmt.Sprintf("&appid=%s", api)

	return aData.Base + aData.Location + aData.ApiKey, nil

}

func GetWeatherAPIKey(filepath string) (string, error) {

	b, err := ioutil.ReadFile(filepath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := strings.TrimSuffix(string(b), "\n")

	return str, nil
}
