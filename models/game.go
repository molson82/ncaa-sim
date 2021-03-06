package models

type Game__c struct {
	Conference string `json:"conference"`
	TeamA      Team   `json:"teamA"`
	TeamB      Team   `json:"teamB"`
	Order      int    `json:"order"`
}
type Game struct {
	GameID                            int           `json:"GameID"`
	Season                            int           `json:"Season"`
	SeasonType                        int           `json:"SeasonType"`
	Status                            string        `json:"Status"`
	Day                               string        `json:"Day"`
	DateTime                          string        `json:"DateTime"`
	AwayTeam                          string        `json:"AwayTeam"`
	HomeTeam                          string        `json:"HomeTeam"`
	AwayTeamID                        int           `json:"AwayTeamID"`
	HomeTeamID                        int           `json:"HomeTeamID"`
	AwayTeamScore                     interface{}   `json:"AwayTeamScore"`
	HomeTeamScore                     interface{}   `json:"HomeTeamScore"`
	Updated                           string        `json:"Updated"`
	Period                            interface{}   `json:"Period"`
	TimeRemainingMinutes              interface{}   `json:"TimeRemainingMinutes"`
	TimeRemainingSeconds              interface{}   `json:"TimeRemainingSeconds"`
	PointSpread                       float64       `json:"PointSpread"`
	OverUnder                         float64       `json:"OverUnder"`
	AwayTeamMoneyLine                 int           `json:"AwayTeamMoneyLine"`
	HomeTeamMoneyLine                 int           `json:"HomeTeamMoneyLine"`
	GlobalGameID                      int           `json:"GlobalGameID"`
	GlobalAwayTeamID                  int           `json:"GlobalAwayTeamID"`
	GlobalHomeTeamID                  int           `json:"GlobalHomeTeamID"`
	TournamentID                      int           `json:"TournamentID"`
	Bracket                           string        `json:"Bracket"`
	Round                             int           `json:"Round"`
	AwayTeamSeed                      int           `json:"AwayTeamSeed"`
	HomeTeamSeed                      int           `json:"HomeTeamSeed"`
	AwayTeamPreviousGameID            int           `json:"AwayTeamPreviousGameID"`
	HomeTeamPreviousGameID            interface{}   `json:"HomeTeamPreviousGameID"`
	AwayTeamPreviousGlobalGameID      int           `json:"AwayTeamPreviousGlobalGameID"`
	HomeTeamPreviousGlobalGameID      interface{}   `json:"HomeTeamPreviousGlobalGameID"`
	TournamentDisplayOrder            int           `json:"TournamentDisplayOrder"`
	TournamentDisplayOrderForHomeTeam string        `json:"TournamentDisplayOrderForHomeTeam"`
	IsClosed                          bool          `json:"IsClosed"`
	GameEndDateTime                   interface{}   `json:"GameEndDateTime"`
	HomeRotationNumber                int           `json:"HomeRotationNumber"`
	AwayRotationNumber                int           `json:"AwayRotationNumber"`
	TopTeamPreviousGameID             interface{}   `json:"TopTeamPreviousGameId"`
	BottomTeamPreviousGameID          int           `json:"BottomTeamPreviousGameId"`
	Channel                           string        `json:"Channel"`
	NeutralVenue                      bool          `json:"NeutralVenue"`
	AwayPointSpreadPayout             int           `json:"AwayPointSpreadPayout"`
	HomePointSpreadPayout             int           `json:"HomePointSpreadPayout"`
	OverPayout                        int           `json:"OverPayout"`
	UnderPayout                       int           `json:"UnderPayout"`
	Stadium                           interface{}   `json:"Stadium"`
	Periods                           []interface{} `json:"Periods"`
}
