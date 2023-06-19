package apimodels

import (
	"time"
)

type MatchAPI struct {
	ID         int64    `json:"id,omitempty"`
	UserAsker  *UserAPI `json:"userAsker,omitempty"`
	UserMatch  *UserAPI `json:"userMatch,omitempty"`
	MatchScore int32    `json:"matchScore,omitempty"`
}

type InterviewAPI struct {
	ID     int64     `json:"id,omitempty"`
	Match  *MatchAPI `json:"match,omitempty"`
	Status string    `json:"status,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}

type RequestToMatchAPI struct {
	ID             int64      `json:"id,omitempty"`
	Skill          *UserAPI   `json:"skill,omitempty"`
	RequiredSkills []SkillAPI `json:"requiredSkills,omitempty"`
}
