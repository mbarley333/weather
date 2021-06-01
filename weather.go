package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// WeatherReponse is used to accept JSON structured
// data.  Output is not very human friendly and is thus
// a stage type for the Weather struct
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

// Weather is the human friendly struct that is populated
// by the WeatherReponse struct
type Weather struct {
	Main        string
	Description string
	Temp        float64
	City        string
}

// Client is used to assemble the necessary parts for a HTTP request
// and includes the HTTPClient lib to set timeout
type Client struct {
	Base       string
	Units      string
	ApiKey     string
	HTTPClient *http.Client
}

// Get takes a location and resturns a Weather struct and error.
func (c Client) Get(location string) (Weather, error) {

	//assemble request based on Client struct data
	url := c.FormatURL(location)

	//use Client since we override default HTTPClient settings -- timeout
	resp, err := c.HTTPClient.Get(url)

	if err != nil {
		return Weather{}, fmt.Errorf("something went wrong.  Please try again later.  %v", err)
	}
	//close when done to prevent resource leaks
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected status code.  %v", resp.StatusCode)
	}

	// pass io.Reader to ParseReponse
	return ParseResponse(resp.Body)

}

// NewClient takes apiKey and tempunits populates the Client struct
// with a base url, temperature units, apikey AND sets up the
// HTTPClient with timeout settings since default timeout is too long.
// Returns Client struct and error
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

// GetWeatherAPIKey takes env as name of environmental variable
// which holds the API key
func GetWeatherAPIKey(env string) (string, error) {

	apikey := os.Getenv(env)

	if len(apikey) < 1 {
		return "", fmt.Errorf("%s value not set", env)
	}
	return apikey, nil
}

// FormatURL is a method on the Client struct that
// assembles the URL used in the request
func (c Client) FormatURL(location string) string {

	return fmt.Sprintf("%s/data/2.5/weather?q=%s&units=%s&appid=%s", c.Base, location, c.Units, c.ApiKey)

}

// ParseReponse takes an io.Reader and decodes into WeatherReponse struct
// The WeatherResonse struct in then used to setup the Weather struct for
// human reading
func ParseResponse(r io.Reader) (Weather, error) {

	var result WeatherResponse

	//decodes io.Reader into variable address (e.g. &result)
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
