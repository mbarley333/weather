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

type WeatherResponse struct {
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
	TempF float64
	TempC float64
}

type Catalog map[string]WeatherResponse

type ApiURL struct {
	Base     string
	Location string
	ApiKey   string
}

func Get(url string) (WeatherResponse, error) {

	resp, err := http.Get(url)

	if err != nil {
		return WeatherResponse{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("unexpected status code.  %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	fmt.Printf("%s", data)
	var wdata WeatherResponse
	err = json.Unmarshal(data, &wdata)
	if err != nil {
		log.Fatal("unable to unmarshall data")
	}

	return wdata, nil
}

func SetApiURL(location string, apiKey string) (string, error) {

	if len(location) == 0 {
		return "", fmt.Errorf("something went wrong.  Please provide location")
	}

	var aData ApiURL

	aData.Base = "https://api.openweathermap.org/data/2.5/weather?q="
	aData.Location = location
	aData.ApiKey = fmt.Sprintf("&appid=%s", apiKey)

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

func (w *WeatherResponse) SetTemp(t float64) {
	w.TempF = (t-273.15)*9/5 + 32
	w.TempC = t - 273.15

}
