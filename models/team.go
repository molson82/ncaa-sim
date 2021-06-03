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
	Statid                    int         `json:"StatID"`
	TeamID                    int         `json:"TeamID"`
	Seasontype                int         `json:"SeasonType"`
	Season                    int         `json:"Season"`
	Name                      string      `json:"Name"`
	Team                      string      `json:"Team"`
	Wins                      int         `json:"Wins"`
	Losses                    int         `json:"Losses"`
	Overtimelosses            int         `json:"OvertimeLosses"`
	Opponentposition          interface{} `json:"OpponentPosition"`
	Globalteamid              int         `json:"GlobalTeamID"`
	Updated                   string      `json:"Updated"`
	Games                     int         `json:"Games"`
	Fantasypoints             float64     `json:"FantasyPoints"`
	Fantasypointsfanduel      float64     `json:"FantasyPointsFanDuel"`
	Fantasypointsdraftkings   float64     `json:"FantasyPointsDraftKings"`
	Fantasypointsyahoo        float64     `json:"FantasyPointsYahoo"`
	Minutes                   int         `json:"Minutes"`
	Seconds                   int         `json:"Seconds"`
	Goals                     float64     `json:"Goals"`
	Assists                   float64     `json:"Assists"`
	Shotsongoal               float64     `json:"ShotsOnGoal"`
	Powerplaygoals            float64     `json:"PowerPlayGoals"`
	Shorthandedgoals          float64     `json:"ShortHandedGoals"`
	Emptynetgoals             float64     `json:"EmptyNetGoals"`
	Powerplayassists          float64     `json:"PowerPlayAssists"`
	Shorthandedassists        float64     `json:"ShortHandedAssists"`
	Hattricks                 float64     `json:"HatTricks"`
	Shootoutgoals             float64     `json:"ShootoutGoals"`
	Plusminus                 float64     `json:"PlusMinus"`
	Penaltyminutes            float64     `json:"PenaltyMinutes"`
	Blocks                    float64     `json:"Blocks"`
	Hits                      float64     `json:"Hits"`
	Takeaways                 float64     `json:"Takeaways"`
	Giveaways                 float64     `json:"Giveaways"`
	Faceoffswon               float64     `json:"FaceoffsWon"`
	Faceoffslost              float64     `json:"FaceoffsLost"`
	Shifts                    float64     `json:"Shifts"`
	Goaltendingminutes        int         `json:"GoaltendingMinutes"`
	Goaltendingseconds        int         `json:"GoaltendingSeconds"`
	Goaltendingshotsagainst   float64     `json:"GoaltendingShotsAgainst"`
	Goaltendinggoalsagainst   float64     `json:"GoaltendingGoalsAgainst"`
	Goaltendingsaves          float64     `json:"GoaltendingSaves"`
	Goaltendingwins           float64     `json:"GoaltendingWins"`
	Goaltendinglosses         float64     `json:"GoaltendingLosses"`
	Goaltendingshutouts       float64     `json:"GoaltendingShutouts"`
	Started                   interface{} `json:"Started"`
	Benchpenaltyminutes       float64     `json:"BenchPenaltyMinutes"`
	Goaltendingovertimelosses float64     `json:"GoaltendingOvertimeLosses"`
	Fantasypointsfantasydraft float64     `json:"FantasyPointsFantasyDraft"`
	Opponentstat              struct {
		Statid                    int         `json:"StatID"`
		Teamid                    int         `json:"TeamID"`
		Seasontype                int         `json:"SeasonType"`
		Season                    int         `json:"Season"`
		Name                      string      `json:"Name"`
		Team                      string      `json:"Team"`
		Wins                      int         `json:"Wins"`
		Losses                    int         `json:"Losses"`
		Overtimelosses            int         `json:"OvertimeLosses"`
		Opponentposition          interface{} `json:"OpponentPosition"`
		Globalteamid              int         `json:"GlobalTeamID"`
		Updated                   string      `json:"Updated"`
		Games                     int         `json:"Games"`
		Fantasypoints             float64     `json:"FantasyPoints"`
		Fantasypointsfanduel      float64     `json:"FantasyPointsFanDuel"`
		Fantasypointsdraftkings   float64     `json:"FantasyPointsDraftKings"`
		Fantasypointsyahoo        float64     `json:"FantasyPointsYahoo"`
		Minutes                   int         `json:"Minutes"`
		Seconds                   int         `json:"Seconds"`
		Goals                     float64     `json:"Goals"`
		Assists                   float64     `json:"Assists"`
		Shotsongoal               float64     `json:"ShotsOnGoal"`
		Powerplaygoals            float64     `json:"PowerPlayGoals"`
		Shorthandedgoals          float64     `json:"ShortHandedGoals"`
		Emptynetgoals             float64     `json:"EmptyNetGoals"`
		Powerplayassists          float64     `json:"PowerPlayAssists"`
		Shorthandedassists        float64     `json:"ShortHandedAssists"`
		Hattricks                 float64     `json:"HatTricks"`
		Shootoutgoals             float64     `json:"ShootoutGoals"`
		Plusminus                 float64     `json:"PlusMinus"`
		Penaltyminutes            float64     `json:"PenaltyMinutes"`
		Blocks                    float64     `json:"Blocks"`
		Hits                      float64     `json:"Hits"`
		Takeaways                 float64     `json:"Takeaways"`
		Giveaways                 float64     `json:"Giveaways"`
		Faceoffswon               float64     `json:"FaceoffsWon"`
		Faceoffslost              float64     `json:"FaceoffsLost"`
		Shifts                    float64     `json:"Shifts"`
		Goaltendingminutes        int         `json:"GoaltendingMinutes"`
		Goaltendingseconds        int         `json:"GoaltendingSeconds"`
		Goaltendingshotsagainst   float64     `json:"GoaltendingShotsAgainst"`
		Goaltendinggoalsagainst   float64     `json:"GoaltendingGoalsAgainst"`
		Goaltendingsaves          float64     `json:"GoaltendingSaves"`
		Goaltendingwins           float64     `json:"GoaltendingWins"`
		Goaltendinglosses         float64     `json:"GoaltendingLosses"`
		Goaltendingshutouts       float64     `json:"GoaltendingShutouts"`
		Started                   interface{} `json:"Started"`
		Benchpenaltyminutes       float64     `json:"BenchPenaltyMinutes"`
		Goaltendingovertimelosses float64     `json:"GoaltendingOvertimeLosses"`
		Fantasypointsfantasydraft float64     `json:"FantasyPointsFantasyDraft"`
		Opponentstat              interface{} `json:"OpponentStat"`
	} `json:"OpponentStat"`
}
