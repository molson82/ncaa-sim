package models

import (
	"encoding/json"
	"fmt"
	"sort"
)

type NCAAMB struct {
	Sport                        string  `json:"sport"`
	LeftTopBracketConference     string  `json:"leftTopBracketConference"`
	LeftBottomBracketConference  string  `json:"leftBottomBracketConference"`
	RightTopBracketConference    string  `json:"rightTopBracketConference"`
	RightBottomBracketConference string  `json:"rightBottomBracketConference"`
	Matches                      []Match `json:"matches"`
}

type NCAAMB__c struct {
	Sport                        string    `json:"sport"`
	LeftTopBracketConference     string    `json:"leftTopBracketConference"`
	LeftBottomBracketConference  string    `json:"leftBottomBracketConference"`
	RightTopBracketConference    string    `json:"rightTopBracketConference"`
	RightBottomBracketConference string    `json:"rightBottomBracketConference"`
	Games                        []Game__c `json:"games"`
}

func ReadInNCAAMBBracket(bracket []byte) (NCAAMB, error) {
	var ncaambObj NCAAMB
	decodeErr := json.Unmarshal(bracket, &ncaambObj)
	if decodeErr != nil {
		return NCAAMB{}, decodeErr
	}
	return ncaambObj, nil
}

func PrintNCAAMBBracket(bracket []byte) error {
	var ncaambObj NCAAMB
	decodeErr := json.Unmarshal(bracket, &ncaambObj)
	if decodeErr != nil {
		return decodeErr
	}
	fmt.Println(ncaambObj)
	return nil
}

func GenerateFinalBracket(teams []Team, bracketObj NCAAMB, teamList map[string]Team__c) (NCAAMB__c, error) {
	var finalBracket NCAAMB__c
	// Initialize constants
	finalBracket.Sport = bracketObj.Sport
	finalBracket.LeftTopBracketConference = bracketObj.LeftTopBracketConference
	finalBracket.LeftBottomBracketConference = bracketObj.LeftBottomBracketConference
	finalBracket.RightTopBracketConference = bracketObj.RightTopBracketConference
	finalBracket.RightBottomBracketConference = bracketObj.RightBottomBracketConference

	var gameList []Game__c

	for _, v := range teamList {
		if doesTeamExist(gameList, v.ID) {
			continue
		}
		// get API team obj
		apiTeam := findAPITeam(teams, v.ID)
		var game Game__c
		game.Conference = bracketObj.Matches[v.Order-1].Conference
		game.Order = v.Order
		if v.AorB == "A" {
			game.TeamA = apiTeam
			game.TeamB = findAPITeam(teams, teamList[bracketObj.Matches[v.Order-1].TeamB].ID)
		} else {
			game.TeamB = apiTeam
			game.TeamA = findAPITeam(teams, teamList[bracketObj.Matches[v.Order-1].TeamA].ID)
		}
		gameList = append(gameList, game)
	}
	finalBracket.Games = gameList
	sort.SliceStable(finalBracket.Games, func(i, j int) bool {
		return finalBracket.Games[i].Order < finalBracket.Games[j].Order
	})

	return finalBracket, nil
}

func doesTeamExist(gamelist []Game__c, id int) bool {
	for _, v := range gamelist {
		if v.TeamA.TeamID == id || v.TeamB.TeamID == id {
			return true
		}
	}
	return false
}

func findAPITeam(teams []Team, id int) Team {
	for _, v := range teams {
		if v.TeamID == id {
			return v
		}
	}
	return Team{}
}
