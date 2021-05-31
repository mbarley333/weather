package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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
	Temp        float64
	City        string
}

type Client struct {
	Base       string
	Units      string
	ApiKey     string
	HTTPClient *http.Client
}

func (c Client) Get(location string) (Weather, error) {

	url := fmt.Sprintf("%s%s%s%s", c.Base, location, c.Units, c.ApiKey)

	//use Client since we override default HTTPClient settings -- timeout
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
	w.Temp = wdata.Main.Temp
	w.City = wdata.City

	//w.SetTemp(wdata.Main.Temp)

	return w, nil
}

func NewClient(apiKey string, tempunits string) (Client, error) {

	var c Client

	if tempunits == "metric" || tempunits == "imperial" || tempunits == "standard" {
		c.Units = "&units=" + tempunits
	}

	c.Base = "https://api.openweathermap.org/data/2.5/weather?q="
	c.ApiKey = fmt.Sprintf("&appid=%s", apiKey)
	//override default timeout settings
	c.HTTPClient = &http.Client{Timeout: 10 * time.Second}

	return c, nil

}

func GetWeatherAPIKey(env string) (string, error) {

	apikey := os.Getenv(env)

	if len(apikey) < 1 {
		return "", fmt.Errorf("%s value not set", env)
	}
	return apikey, nil
}
