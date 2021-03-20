package models

type Tournament struct {
	TournamentID                 int    `json:"TournamentID"`
	Season                       int    `json:"Season"`
	Name                         string `json:"Name"`
	Location                     string `json:"Location"`
	LeftTopBracketConference     string `json:"LeftTopBracketConference"`
	LeftBottomBracketConference  string `json:"LeftBottomBracketConference"`
	RightTopBracketConference    string `json:"RightTopBracketConference"`
	RightBottomBracketConference string `json:"RightBottomBracketConference"`
	Games                        []Game `json:"Games"`
}
