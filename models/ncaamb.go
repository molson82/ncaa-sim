package models

import (
	"encoding/json"
	"fmt"
)

type NCAAMB struct {
	Sport                        string  `json:"sport"`
	LeftTopBracketConference     string  `json:"leftTopBracketConference"`
	LeftBottomBracketConference  string  `json:"leftBottomBracketConference"`
	RightTopBracketConference    string  `json:"rightTopBracketConference"`
	RightBottomBracketConference string  `json:"rightBottomBracketConference"`
	Matches                      []Match `json:"matches"`
}

func ReadInNCAAMBBracket(bracket []byte) (NCAAMB, error) {
	var ncaambObj NCAAMB
	decodeErr := json.Unmarshal(bracket, &ncaambObj)
	if decodeErr != nil {
		return NCAAMB{}, decodeErr
	}
	return ncaambObj, nil
}

func PrintNCAAMBBracket(bracket []byte) error {
	var ncaambObj NCAAMB
	decodeErr := json.Unmarshal(bracket, &ncaambObj)
	if decodeErr != nil {
		return decodeErr
	}
	fmt.Println(ncaambObj)
	return nil
}
