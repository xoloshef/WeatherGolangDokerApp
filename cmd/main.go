package main

import (
	"WeatherGoDokerApp/internal/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
