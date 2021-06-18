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

## Future enhancements
* output in human friendly format (e.g. Your current weather is clear skies, wind at 10 mph and a tempurature of 78.0 ºF)
* add more data elements to output
* use weather client data as part of a UI project
* use weather client as part of a HTTP server project that uses concurrency


