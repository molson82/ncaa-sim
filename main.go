package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/molson82/ncaa-sim/controllers"
	"github.com/molson82/ncaa-sim/models"
)

var api_key string
var format string

//go:embed "ncaamb_bracket.json"
var ncaamb_bracket []byte

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
	// set globals
	api_key = os.Getenv("API_KEY")
	format = "json"
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}

func main() {
	initEnv()

	season, err1 := controllers.GetCurrentSeason(api_key, format)
	handleError(err1)
	fmt.Printf("Current Season: %v\n", season.Season)

	fmt.Println("Getting All Team Stats...")
	teams, err2 := controllers.GetTeamStats(fmt.Sprint(season.Season), api_key, format)
	handleError(err2)

	fmt.Printf("Done.\nReading in Bracket...\n")
	ncaambObj, err3 := models.ReadInNCAAMBBracket(ncaamb_bracket)
	handleError(err3)

	fmt.Printf("Done.\nFinding Teams In Bracket...\n")
	teamList := findNCAAMBTeams(teams, ncaambObj)
	for k, v := range teamList {
		fmt.Printf("\n%v = %v | W: %v, L: %v\n", k, v.Name, v.Wins, v.Losses)
	}

	// marchMad, err2 := controllers.GetTournamentHierarchy(season.Season, api_key, format)
	// if err2 != nil {
	// 	fmt.Printf("ERROR: %v\n", err2)
	// }
	// fmt.Printf("Tournament Name: %v\n", marchMad.Name)
}
