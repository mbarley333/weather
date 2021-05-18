package main

import (
	"log"
	"weather"
)

func main() {

	data, err := weather.Get("Kaneohe")
	if err != nil {
		log.Fatal("error")
	}

	//refactor to call get

	//pass json into something human friendly

	//struct summary clouds, temp

	//unit tests for offline testing

	//url config to point to local http listen

	//fake api

	//given location, construct the weathermap url

	//pass json into go struct

	//unit test ^^^

	//construct url

	//decode json

}
