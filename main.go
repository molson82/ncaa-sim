package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/molson82/ncaa-sim/controllers"
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

	season, err1 := controllers.GetCurrentSeason(api_key, format)
	if err1 != nil {
		fmt.Printf("ERROR: %v\n", err1)
	}
	fmt.Printf("Current Season: %v\n", season.Season)

	marchMad, err2 := controllers.GetTournamentHierarchy(season.Season, api_key, format)
	if err2 != nil {
		fmt.Printf("ERROR: %v\n", err2)
	}
	fmt.Printf("Tournament Name: %v\n", marchMad.Name)
	fmt.Printf("Number of Games %v\n", len(marchMad.Games))
	controllers.PrintTournamentHierarchyGames(marchMad.Games)
}
