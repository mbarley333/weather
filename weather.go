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
	City string `json:"name"`
}

type Weather struct {
	Main        string
	Description string
	TempK       float64
	TempF       float64
	TempC       float64
	City        string
}

type Client struct {
	Base       string
	ApiKey     string
	HTTPClient *http.Client
}

func (c Client) Get(location string) (Weather, error) {

	url := fmt.Sprintf("%s%s%s", c.Base, location, c.ApiKey)

	resp, err := c.HTTPClient.Get(url)

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

	w.SetTemp(wdata.Main.Temp)

	return w, nil
}

func NewClient(apiKey string) (Client, error) {

	var c Client

	c.Base = "https://api.openweathermap.org/data/2.5/weather?q="
	c.ApiKey = fmt.Sprintf("&appid=%s", apiKey)
	c.HTTPClient = &http.Client{}

	return c, nil

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
