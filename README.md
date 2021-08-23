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



