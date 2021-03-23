package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/molson82/ncaa-sim/models"
)

type team struct {
	ID     int
	Name   string
	Team   string
	Wins   int
	Losses int
}

func findNCAAMBTeams(apiTeams []models.Team, bracket models.NCAAMB) map[string]team {
	results := make(map[string]team)
	duplicates := make(map[string][]team)
	for _, v := range bracket.Matches {
		var count int
		for _, q := range apiTeams {
			if (strings.Contains(q.Name, v.TeamA)) || (strings.Contains(q.Team, v.TeamA)) {
				if count > 0 {
					duplicates[v.TeamA] = append(duplicates[v.TeamA], team{q.TeamID, q.Name, q.Team, q.Wins, q.Losses})
					continue
				}
				results[v.TeamA] = team{q.TeamID, q.Name, q.Team, q.Wins, q.Losses}
				count++
			}
		}
	}
	return handleDuplicates(results, duplicates)
}

func handleDuplicates(result map[string]team, dups map[string][]team) map[string]team {
	finalResults := make(map[string]team)
OUTER:
	for k, v := range dups {
		fmt.Printf("\n-----------------------------------------------------------------------------")
		fmt.Printf("\n%v Results for: %v\n[", len(v)+1, k)
		fmt.Printf("%v, ", result[k].Name)
		for _, j := range v {
			fmt.Printf("%v, ", j.Name)
		}
		fmt.Printf("]\n")

		c := askForConfirmation(fmt.Sprintf("\nIs this the correct one?\nName: %v | Team: %v | Wins: %v | Losses: %v  ", result[k].Name, result[k].Team, result[k].Wins, result[k].Losses))
		if c {
			finalResults[k] = result[k]
			continue
		}
		for _, j := range v {
			q := askForConfirmation(fmt.Sprintf("\nIs this the correct one?\nName: %v | Team: %v | Wins: %v | Losses: %v  ", j.Name, j.Team, j.Wins, j.Losses))
			if q {
				finalResults[k] = j
				continue OUTER
			}
		}
	}

	return finalResults
}

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
