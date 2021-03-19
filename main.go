package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/molson82/ncaa-sim/models"
)

var api_key string
var format string

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
	// set globals
	api_key = os.Getenv("API_KEY")
	format = "json"
}

func main() {
	initEnv()

	fmt.Println("Getting current NCAA Season...")
	resp, reqErr := http.Get(fmt.Sprintf("https://api.sportsdata.io/v3/cbb/scores/json/CurrentSeason?key=%s&format=%s", api_key, format))
	if reqErr != nil {
		fmt.Printf("ERROR: %v\n", reqErr)
	}

	var season models.Season
	decodeErr := render.DecodeJSON(resp.Body, &season)
	if decodeErr != nil {
		fmt.Printf("ERROR: %v\n", decodeErr)
	}
	fmt.Printf("Current Season: %v\n", season.Season)

	fmt.Printf("Getting Tournament for Season: %v\n", season.Season)
	resp2, reqErr2 := http.Get(fmt.Sprintf("https://api.sportsdata.io/v3/cbb/scores/json/Tournament/2021?key=%s&format=%s", api_key, format))
	if reqErr2 != nil {
		fmt.Printf("ERROR: %v\n", reqErr2)
	}

	var marchMad models.Tournament
	decodeErr2 := render.DecodeJSON(resp2.Body, &marchMad)
	if decodeErr2 != nil {
		fmt.Printf("ERROR: %v\n", decodeErr2)
	}
	fmt.Printf("Tournament Name: %v\n", marchMad.Name)
	fmt.Printf("Number of Games %v\n", len(marchMad.Games)+1)
}
