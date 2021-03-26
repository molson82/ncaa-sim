package main

import (
	"fmt"
	"math"

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
	teamAScore := calculateTeamScore_average(teamA)
	teamBScore := calculateTeamScore_average(teamB)

	fmt.Printf("\n%v Score: %v | %v Score: %v\n", teamA.Name, teamAScore, teamB.Name, teamBScore)

	if teamAScore >= teamBScore {
		return teamA
	}
	return teamB
}

func calculateTeamScore_average(t models.Team) float64 {
	return (calculateTeamScore_v1(t) + calculateTeamScore_v2(t) +
		calculateTeamScore_v3(t) + calculateTeamScore_v4(t)) / 4
}

func calculateTeamScore_v1(t models.Team) float64 {
	p1 := math.Abs(float64(t.BlockedShots) - float64(t.Steals))
	p2 := ((t.OffensiveRebounds + t.DefensiveRebounds) - (t.PersonalFouls + t.Turnovers))
	p3 := p1 * float64(p2)
	p4 := float64(t.ThreePointersAttempted) * t.ThreePointersPercentage
	p5 := float64(t.TwoPointersAttempted) * t.TwoPointersPercentage
	p6 := float64(t.FreeThrowsAttempted) * t.FreeThrowsPercentage
	p7 := p3 + p4 + p5 + p6
	p8 := float64(t.Points) / float64(t.Games)
	p9 := float64(t.FantasyPoints) / p8
	p10 := p7 + p9

	return p10
}

func calculateTeamScore_v2(t models.Team) float64 {
	p1 := math.Abs(float64(t.BlockedShots) - float64(t.Steals))
	p2 := ((t.OffensiveRebounds + t.DefensiveRebounds) - (t.PersonalFouls + t.Turnovers))
	p3 := p1 * float64(p2)
	p4 := float64(t.ThreePointersAttempted) * t.ThreePointersPercentage
	p5 := float64(t.TwoPointersAttempted) * t.TwoPointersPercentage
	p6 := float64(t.FreeThrowsAttempted) * t.FreeThrowsPercentage
	p7 := p3 + p4 + p5 + p6
	p8 := float64(t.Points) / float64(t.Games)
	p9 := float64(t.FantasyPoints) / p8
	p10 := p7 / p9

	return p10
}

func calculateTeamScore_v3(t models.Team) float64 {
	p1 := math.Abs(float64(t.BlockedShots) - float64(t.Steals))
	p2 := ((t.OffensiveRebounds + t.DefensiveRebounds) - (t.PersonalFouls + t.Turnovers))
	p3 := p1 * float64(p2)
	p4 := float64(t.ThreePointersAttempted) * t.ThreePointersPercentage
	p5 := float64(t.TwoPointersAttempted) * t.TwoPointersPercentage
	p6 := float64(t.FreeThrowsAttempted) * t.FreeThrowsPercentage
	p7 := p3 + p4 + p5 + p6
	p8 := float64(t.Points) / float64(t.Games)
	p9 := float64(t.FantasyPoints) / p8
	p10 := p7 + p9
	p11 := p7 / p9
	p12 := (p10 - p11) * math.Pi

	return p12
}

func calculateTeamScore_v4(t models.Team) float64 {
	p1 := math.Abs(float64(t.BlockedShots) - float64(t.Steals))
	p2 := ((t.OffensiveRebounds + t.DefensiveRebounds) - (t.PersonalFouls + t.Turnovers))
	p3 := p1 * float64(p2)
	p4 := float64(t.ThreePointersAttempted) * t.ThreePointersPercentage
	p5 := float64(t.TwoPointersAttempted) * t.TwoPointersPercentage
	p6 := float64(t.FreeThrowsAttempted) * t.FreeThrowsPercentage
	p7 := p3 + p4 + p5 + p6
	p8 := float64(t.Points) / float64(t.Games)
	p9 := float64(t.FantasyPoints) / p8
	p10 := p7 / p9

	pdifShots := ((math.Abs(p4 - p6)) / ((p4 + p6) / 2)) * 100
	pdifShots2 := ((math.Abs(p4 - p5)) / ((p4 + p5) / 2)) * 100
	p11 := ((math.Abs(pdifShots - pdifShots2)) / ((pdifShots + pdifShots2) / 2)) * 100
	p12 := math.Mod((p10*p11), math.Pi) * math.Phi

	return p12
}
