package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/molson82/ncaa-sim/models"
)

func findNCAAMBTeams(apiTeams []models.Team, bracket models.NCAAMB) map[string]models.Team__c {
	results := make(map[string]models.Team__c)
	duplicates := make(map[string][]models.Team__c)
	for _, v := range bracket.Matches {
		var count int
		for _, q := range apiTeams {
			if (strings.Contains(strings.ToLower(q.Name), strings.ToLower(v.TeamA))) ||
				(strings.Contains(strings.ToLower(q.Team), strings.ToLower(v.TeamA))) {
				if count > 0 {
					duplicates[v.TeamA] = append(duplicates[v.TeamA], models.Team__c{q.TeamID, q.Name, q.Team, q.Wins, q.Losses, v.Order, "A"})
					continue
				}
				results[v.TeamA] = models.Team__c{q.TeamID, q.Name, q.Team, q.Wins, q.Losses, v.Order, "A"}
				count++
			}
		}
	}
	return handleDuplicates(results, duplicates)
}

func handleDuplicates(result map[string]models.Team__c, dups map[string][]models.Team__c) map[string]models.Team__c {
	finalResults := result
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
