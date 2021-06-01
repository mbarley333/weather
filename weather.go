package weather

import (
	"encoding/json"
	"fmt"
	"io"
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

	url := c.FormatURL(location)
	//use Client since we override default HTTPClient settings -- timeout
	resp, err := c.HTTPClient.Get(url)

	if err != nil {
		return Weather{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected status code.  %v", resp.StatusCode)
	}

	return ParseResponse(resp.Body)

}

func NewClient(apiKey string, tempunits string) (Client, error) {

	var c Client

	if tempunits == "metric" || tempunits == "imperial" || tempunits == "standard" {
		c.Units = tempunits
	}

	c.Base = "https://api.openweathermap.org"
	c.ApiKey = apiKey
	//override default HTTPClient timeout settings
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

func (c Client) FormatURL(location string) string {

	return fmt.Sprintf("%s/data/2.5/weather?q=%s&units=%s&appid=%s", c.Base, location, c.Units, c.ApiKey)

}

func ParseResponse(r io.Reader) (Weather, error) {

	var result WeatherResponse
	err := json.NewDecoder(r).Decode(&result)
	if err != nil {
		return Weather{}, err
	}

	var w Weather

	w.Main = result.Weather[0].Main
	w.Description = result.Weather[0].Description
	w.Temp = result.Main.Temp
	w.City = result.City
	return w, nil
}
