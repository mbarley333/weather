# Weather Client

A simple command line utility to query the OpenWeatherMap Current Weather Data API

```bash
./weather -location Seattle -units imperial
{Clouds few clouds 81.86 Seattle}
```

## Usage
* Prior to use, install [Golang](https://golang.org/doc/install)
* Create an account on [Open Weather Map](https://home.openweathermap.org/users/sign_up) and sign up for an [API key](https://home.openweathermap.org/api_keys)
* Create an environment variable for your Open Weather Map API key: `export WEATHERAPI=YourOpenWeatherMapAPIKey`
* Clone weather repo to local machine and change to that directory
* Build weather client: go build -o weather ./cmd/main.go
* ./weather -location Seattle -units imperial


## Goals
To learn and become more familiar with the following aspects of the Go language:
* testing
* functions and methods
* structs
* HTTP Client
* writing clear code


## Application Requirements
* I want to be able to connect an external REST API -- OpenWeatherAPI
* I want to be able to pass an formatted URL to the REST API
* I want to take the json response from the API and populate a response struct
* I want to use the reponse struct to populate a final struct for display to the user


## Testing Requirements
* I want to test the behavior of the application code as it relates to an external API and the Go standard library
* I don't want to test the OpenWeather API since I assume they test their code
* I don't want to test the Go standard library since they do test their code


## Tests
* TestFormatURL: Given a set of input parameters, can the FormatURL method correctly put together a url string
* TestGetWeather: Given a test web sever w/ data, can the Get method correctly return a Weather struct with the fields populated
* TestNewClient: Given a set of input parameters, can the NewClient func correctly return a Client struct with the fields populated
* TestUnmarshallJson: Given an io.Reader, can the ParseResponse func correctly decode json to a Weather struct


## Structs and Methods
Client: Holds all info needed to make a call to the OpenWeather API
* FormatURL: Builds a URL based off of input parameters.  URL will be used in call to open OpenWeather API
* Get: Makes the call to OpenWeatherAPI, converts JSON to struct for use in customer output

WeatherResponse: Populated from OpenWeather API JSON response.  Not fit for output to customer

Weather: Populated from WeatherResponse struct.  Used to present data to customer


## Standalone Funcs
NewClient: Generates a Client struct for use in querying the OpenWeather API

ParseResponse: Accepts an io.Reader and decodes JSON into WeatherResponse struct.  WeatherResponse struct in then used to populate a Weather struct.
