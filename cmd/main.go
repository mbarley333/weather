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

	//refactor to call get -- done

	//pass json into something human friendly -- done

	//struct summary clouds, temp -- done

	//unit tests for offline testing -- done

	//url config to point to local http listen

}
