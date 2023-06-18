package models

import (
	"time"
)

type MatchHandler struct {
	ID         int64        `json:"id,omitempty"`
	UserAsker  *UserHandler `json:"userAsker,omitempty"`
	UserMatch  *UserHandler `json:"userMatch,omitempty"`
	MatchScore int32        `json:"matchScore,omitempty"`
}

type InterviewHandler struct {
	ID     int64         `json:"id,omitempty"`
	Match  *MatchHandler `json:"match,omitempty"`
	Status string        `json:"status,omitempty"`
	Date   time.Time     `json:"date,omitempty"`
}

type RequestToMatchHandler struct {
	ID             int64          `json:"id,omitempty"`
	Skill          *UserHandler   `json:"skill,omitempty"`
	RequiredSkills []SkillHandler `json:"requiredSkills,omitempty"`
}
