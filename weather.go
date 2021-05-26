package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	} `json:"coord"`
}

type Weather struct {
	Main        string
	Description string
	TempK       float64
	TempF       float64
	TempC       float64
	City        string
	Latitude    float64
	Longitude   float64
}

type ApiURL struct {
	Base     string
	Location string
	ApiKey   string
}

func Get(url string) (Weather, error) {

	resp, err := http.Get(url)

	if err != nil {
		return Weather{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected status code.  %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}

	var wdata WeatherResponse
	err = json.Unmarshal(data, &wdata)
	if err != nil {
		log.Fatal("unable to unmarshall data")
	}

	var w Weather

	w.Main = wdata.Weather[0].Main
	w.Description = wdata.Weather[0].Description
	w.TempK = wdata.Main.Temp
	w.City = wdata.City
	w.Latitude = wdata.Coord.Latitude
	w.Longitude = wdata.Coord.Longitude

	return w, nil
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

func GetWeatherAPIKey(env string) (string, error) {

	apikey := os.Getenv(env)

	if len(apikey) < 1 {
		return "", fmt.Errorf("%s value not set", env)
	}
	return apikey, nil
}

func (w *Weather) SetTemp(t float64) {

	w.TempF = (t-273.15)*9/5 + 32
	w.TempC = t - 273.15

}
