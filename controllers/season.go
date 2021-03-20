package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/molson82/ncaa-sim/models"
)

func GetCurrentSeason(key string, format string) (models.Season, error) {
	fmt.Println("Getting current NCAA Season...")
	resp, reqErr := http.Get(fmt.Sprintf("https://api.sportsdata.io/v3/cbb/scores/json/CurrentSeason?key=%s&format=%s", key, format))
	if reqErr != nil {
		return models.Season{}, reqErr
	}

	var season models.Season
	decodeErr := render.DecodeJSON(resp.Body, &season)
	if decodeErr != nil {
		return models.Season{}, decodeErr
	}

	return season, nil
}
