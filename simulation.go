package main

import "github.com/molson82/ncaa-sim/models"

func simulateRound(bracket models.NCAAMB__c) (models.NCAAMB__c, error) {
	var newBracket models.NCAAMB__c
	newBracket.Sport = bracket.Sport
	newBracket.LeftTopBracketConference = bracket.LeftTopBracketConference
	newBracket.LeftBottomBracketConference = bracket.LeftBottomBracketConference
	newBracket.RightTopBracketConference = bracket.RightTopBracketConference
	newBracket.RightBottomBracketConference = bracket.RightBottomBracketConference

	var newGames []models.Game__c
	j := 1
	for i := 0; i < len(bracket.Games)-1; i += 2 {
		var newGame models.Game__c
		newGame.Conference = bracket.Games[i].Conference
		newGame.TeamA = findWinner(bracket.Games[i].TeamA, bracket.Games[i].TeamB)
		newGame.TeamB = findWinner(bracket.Games[i+1].TeamA, bracket.Games[i+1].TeamB)
		newGame.Order = j
		newGames = append(newGames, newGame)
		j++
	}
	newBracket.Games = newGames

	return newBracket, nil
}

func findWinner(teamA models.Team, teamB models.Team) models.Team {
	if teamA.Wins >= teamB.Wins {
		return teamA
	}
	return teamB
}
