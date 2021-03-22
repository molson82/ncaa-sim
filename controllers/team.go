package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/molson82/ncaa-sim/models"
)

func GetTeamStats(season string, key string, format string) ([]models.Team, error) {
	resp, reqErr := http.Get(fmt.Sprintf("https://api.sportsdata.io/v3/cbb/scores/json/TeamSeasonStats/%s?key=%s&format=%s", season, key, format))
	if reqErr != nil {
		return []models.Team{}, reqErr
	}

	var teams []models.Team
	decodeErr := render.DecodeJSON(resp.Body, &teams)
	if decodeErr != nil {
		return []models.Team{}, decodeErr
	}

	return teams, nil
}
