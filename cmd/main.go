package main

import (
	"flag"
	"fmt"
	"log"
	"weather"
)

func main() {

	//flags package
	locationPtr := flag.String("location", "Kaneohe", "location name for weather (e.g. Kaneohe")

	//could be friendlier unit names
	tempUnitsPtr := flag.String("units", "imperial", "temparature units of measurement")

	flag.Parse()
	apiKey, err := weather.GetWeatherAPIKey("WEATHERAPI")
	if err != nil {
		log.Fatal("Unable to get API key")
	}

	client, err := weather.NewClient(apiKey, *tempUnitsPtr)
	if err != nil {
		log.Fatal("Something went wrong")
	}

	weather, err := client.Get(*locationPtr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weather)

	//refactor to call get -- done

	//pass json into something human friendly -- done

	//struct summary clouds, temp -- done

	//unit tests for offline testing -- done

	//url config to point to local http listen -- done

	//how would you build open weather map api?

	//now that i can get a response can i pass this to a ui?

}
