package model

import (
	skill "github.com/porky256/mock-interview-tbr/internal/skill/model"
	user "github.com/porky256/mock-interview-tbr/internal/user/model"
	"time"
)

type Match struct {
	Id         int64      `json:"id,omitempty"`
	UserAsker  *user.User `json:"userAsker,omitempty"`
	UserMatch  *user.User `json:"userMatch,omitempty"`
	MatchScore int32      `json:"matchScore,omitempty"`
}

type Interview struct {
	Id     int64     `json:"id,omitempty"`
	Match  *Match    `json:"match,omitempty"`
	Status string    `json:"status,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}

type RequestToMatch struct {
	Id             int64         `json:"id,omitempty"`
	Skill          *user.User    `json:"skill,omitempty"`
	RequiredSkills []skill.Skill `json:"requiredSkills,omitempty"`
}
