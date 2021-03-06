package models

type Team__c struct {
	ID     int
	Name   string
	Team   string
	Wins   int
	Losses int
	Order  int
	AorB   string
}

type Team struct {
	StatID                        int         `json:"StatID"`
	TeamID                        int         `json:"TeamID"`
	SeasonType                    int         `json:"SeasonType"`
	Season                        int         `json:"Season"`
	Name                          string      `json:"Name"`
	Team                          string      `json:"Team"`
	Wins                          int         `json:"Wins"`
	Losses                        int         `json:"Losses"`
	ConferenceWins                int         `json:"ConferenceWins"`
	ConferenceLosses              int         `json:"ConferenceLosses"`
	GlobalTeamID                  int         `json:"GlobalTeamID"`
	Possessions                   float64     `json:"Possessions"`
	Updated                       string      `json:"Updated"`
	Games                         int         `json:"Games"`
	FantasyPoints                 float64     `json:"FantasyPoints"`
	Minutes                       int         `json:"Minutes"`
	FieldGoalsMade                int         `json:"FieldGoalsMade"`
	FieldGoalsAttempted           int         `json:"FieldGoalsAttempted"`
	FieldGoalsPercentage          float64     `json:"FieldGoalsPercentage"`
	EffectiveFieldGoalsPercentage float64     `json:"EffectiveFieldGoalsPercentage"`
	TwoPointersMade               int         `json:"TwoPointersMade"`
	TwoPointersAttempted          int         `json:"TwoPointersAttempted"`
	TwoPointersPercentage         float64     `json:"TwoPointersPercentage"`
	ThreePointersMade             int         `json:"ThreePointersMade"`
	ThreePointersAttempted        int         `json:"ThreePointersAttempted"`
	ThreePointersPercentage       float64     `json:"ThreePointersPercentage"`
	FreeThrowsMade                int         `json:"FreeThrowsMade"`
	FreeThrowsAttempted           int         `json:"FreeThrowsAttempted"`
	FreeThrowsPercentage          float64     `json:"FreeThrowsPercentage"`
	OffensiveRebounds             int         `json:"OffensiveRebounds"`
	DefensiveRebounds             int         `json:"DefensiveRebounds"`
	Rebounds                      int         `json:"Rebounds"`
	OffensiveReboundsPercentage   interface{} `json:"OffensiveReboundsPercentage"`
	DefensiveReboundsPercentage   interface{} `json:"DefensiveReboundsPercentage"`
	TotalReboundsPercentage       interface{} `json:"TotalReboundsPercentage"`
	Assists                       int         `json:"Assists"`
	Steals                        int         `json:"Steals"`
	BlockedShots                  int         `json:"BlockedShots"`
	Turnovers                     int         `json:"Turnovers"`
	PersonalFouls                 int         `json:"PersonalFouls"`
	Points                        int         `json:"Points"`
	TrueShootingAttempts          float64     `json:"TrueShootingAttempts"`
	TrueShootingPercentage        float64     `json:"TrueShootingPercentage"`
	PlayerEfficiencyRating        interface{} `json:"PlayerEfficiencyRating"`
	AssistsPercentage             interface{} `json:"AssistsPercentage"`
	StealsPercentage              interface{} `json:"StealsPercentage"`
	BlocksPercentage              interface{} `json:"BlocksPercentage"`
	TurnOversPercentage           interface{} `json:"TurnOversPercentage"`
	UsageRatePercentage           interface{} `json:"UsageRatePercentage"`
	FantasyPointsFanDuel          float64     `json:"FantasyPointsFanDuel"`
	FantasyPointsDraftKings       float64     `json:"FantasyPointsDraftKings"`
	FantasyPointsYahoo            interface{} `json:"FantasyPointsYahoo"`
}
