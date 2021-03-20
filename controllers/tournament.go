package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/molson82/ncaa-sim/models"
)

func GetTournamentHierarchy(season int, key string, format string) (models.Tournament, error) {
	fmt.Printf("Getting Tournament for Season: %v\n", season)
	resp, reqErr := http.Get(fmt.Sprintf("https://api.sportsdata.io/v3/cbb/scores/json/Tournament/%v?key=%s&format=%s", season, key, format))
	if reqErr != nil {
		return models.Tournament{}, reqErr
	}

	var marchMad models.Tournament
	decodeErr := render.DecodeJSON(resp.Body, &marchMad)
	if decodeErr != nil {
		return models.Tournament{}, decodeErr
	}

	return marchMad, nil
}

func PrintTournamentHierarchyGames(games []models.Game) {
	for _, v := range games {
		fmt.Println("--------------------------------------")
		fmt.Printf("Game Day: %v\n", v.Day)
		fmt.Printf("Away Team: %v\n", v.AwayTeam)
		fmt.Printf("Home Day: %v\n", v.HomeTeam)
		fmt.Printf("Bracket: %v\n", v.Bracket)
		fmt.Printf("Round: %v\n", v.Round)
	}
}
