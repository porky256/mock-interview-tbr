package handlermodel

import (
	"time"
)

type Match struct {
	ID         int64 `json:"id,omitempty"`
	UserAsker  *User `json:"userAsker,omitempty"`
	UserMatch  *User `json:"userMatch,omitempty"`
	MatchScore int32 `json:"matchScore,omitempty"`
}

type Interview struct {
	ID     int64     `json:"id,omitempty"`
	Match  *Match    `json:"match,omitempty"`
	Status string    `json:"status,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}

type RequestToMatch struct {
	ID             int64   `json:"id,omitempty"`
	Skill          *User   `json:"skill,omitempty"`
	RequiredSkills []Skill `json:"requiredSkills,omitempty"`
}
