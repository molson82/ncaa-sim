package models

type Season struct {
	Season                 int    `json:"Season"`
	StartYear              int    `json:"StartYear"`
	EndYear                int    `json:"EndYear"`
	Description            string `json:"Description"`
	RegularSeasonStartDate string `json:"RegularSeasonStartDate"`
	PostSeasonStartDate    string `json:"PostSeasonStartDate"`
	APISeason              string `json:"ApiSeason"`
}
