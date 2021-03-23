package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/molson82/ncaa-sim/controllers"
	"github.com/molson82/ncaa-sim/models"
)

var api_key string
var format string

//go:embed "ncaamb_test.json"
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
	// fmt.Printf("Done.\nFinal List of Teams: \n")
	// for k, v := range teamList {
	// 	fmt.Printf("\n%v = %v | W: %v, L: %v\n", k, v.Name, v.Wins, v.Losses)
	// }

	fmt.Printf("\nBulding Final Bracket with new team selections...\n")
	finalBracket, err4 := models.GenerateFinalBracket(teams, ncaambObj, teamList)
	handleError(err4)
	// for _, v := range finalBracket.Games {
	// 	fmt.Printf("\n%v\nvs\n%v\n", v.TeamA.Name, v.TeamB.Name)
	// }
	fmt.Printf("\nSimulating First Round...")
	round1Bracket, err5 := simulateRound(finalBracket)
	handleError(err5)
	// for _, v := range round1Bracket.Games {
	// 	fmt.Printf("\n%v\nvs\n%v\n", v.TeamA.Name, v.TeamB.Name)
	// }

	fmt.Printf("\nSimulating Second Round...")
	round2Bracket, err7 := simulateRound(round1Bracket)
	handleError(err7)
	// for _, v := range round2Bracket.Games {
	// 	fmt.Printf("\n%v\nvs\n%v\n", v.TeamA.Name, v.TeamB.Name)
	// }

	fmt.Printf("\nSimulating Third Round...")
	round3Bracket, err8 := simulateRound(round2Bracket)
	handleError(err8)
	// for _, v := range round3Bracket.Games {
	// 	fmt.Printf("\n%v\nvs\n%v\n", v.TeamA.Name, v.TeamB.Name)
	// }

	fmt.Printf("\nSimulating Fourth Round...")
	round4Bracket, err9 := simulateRound(round3Bracket)
	handleError(err9)
	// for _, v := range round4Bracket.Games {
	// 	fmt.Printf("\n%v\nvs\n%v\n", v.TeamA.Name, v.TeamB.Name)
	// }

	fmt.Printf("\nSimulating Fifth Round...")
	round5Bracket, err10 := simulateRound(round4Bracket)
	handleError(err10)
	for _, v := range round5Bracket.Games {
		fmt.Printf("\n%v\nvs\n%v\n", v.TeamA.Name, v.TeamB.Name)
	}

	fmt.Printf("\nSimulating Winning Round...")
	winner := findWinner(round5Bracket.Games[0].TeamA, round5Bracket.Games[0].TeamB)
	fmt.Printf("\n\nWinner is %v\n", winner.Name)

	fmt.Printf("\nStarting web server...\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/header.tmpl.html", "templates/footer.tmpl.html"))
		tmpl.Execute(w, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("server is running on : http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
