package main

import (
	"fmt"

	"github.com/molson82/ncaa-sim/models"
)

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
	teamAScore := calculateTeamScore_v1(teamA)
	teamBScore := calculateTeamScore_v1(teamB)

	fmt.Printf("\n%v Score: %v | %v Score: %v\n", teamA.Name, teamAScore, teamB.Name, teamBScore)

	if teamAScore >= teamBScore {
		return teamA
	}
	return teamB
}

func calculateTeamScore_v1(t models.Team) float64 {
	//p1 := math.Abs(float64(t.BlockedShots) - float64(t.Steals))
	//p2 := ((t.OffensiveRebounds + t.DefensiveRebounds) - (t.PersonalFouls + t.Turnovers))
	//p3 := p1 * float64(p2)
	//p4 := float64(t.ThreePointersAttempted) * t.ThreePointersPercentage
	//p5 := float64(t.TwoPointersAttempted) * t.TwoPointersPercentage
	//p6 := float64(t.FreeThrowsAttempted) * t.FreeThrowsPercentage
	//p7 := p3 + p4 + p5 + p6
	//p8 := float64(t.Points) / float64(t.Games)
	//p9 := float64(t.FantasyPoints) / p8
	//p10 := p7 + p9

	return float64(t.Fantasypointsfanduel)
}
