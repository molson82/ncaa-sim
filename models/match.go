package models

type Match struct {
	Conference string `json:"conference"`
	TeamA      string `json:"teamA"`
	TeamB      string `json:"teamB"`
	Order      int    `json:"order"`
}
