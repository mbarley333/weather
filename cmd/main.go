package main

import (
	"flag"
	"fmt"
	"log"
	"weather"
)

func main() {

	apiKey, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		log.Fatal("Unable to get API key")
	}
	client, err := weather.NewClient(apiKey)
	if err != nil {
		log.Fatal("Something went wrong")
	}
	//flags package
	locationPtr := flag.String("location", "foo", "a string")

	flag.Parse()

	weather, err := client.Get(*locationPtr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weather)

	//refactor to call get -- done

	//pass json into something human friendly -- done

	//struct summary clouds, temp -- done

	//unit tests for offline testing -- done

	//url config to point to local http listen

}
