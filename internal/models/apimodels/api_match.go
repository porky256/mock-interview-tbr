package apimodels

import (
	"time"
)

type MatchAPI struct {
	ID         int64    `json:"id,omitempty" validate:"required"`
	UserAsker  *UserAPI `json:"userAsker,omitempty" validate:"required"`
	UserMatch  *UserAPI `json:"userMatch,omitempty" validate:"required"`
	MatchScore int32    `json:"matchScore,omitempty" validate:"required"`
}

type InterviewAPI struct {
	ID     int64     `json:"id,omitempty" validate:"required"`
	Match  *MatchAPI `json:"match,omitempty" validate:"required"`
	Status Status    `json:"status,omitempty" validate:"required, status_custom_validation"`
	Date   time.Time `json:"date,omitempty" validate:"required"`
}

type RequestToMatchAPI struct {
	ID             int64      `json:"id,omitempty" validate:"required"`
	User           *UserAPI   `json:"skill,omitempty" validate:"required"`
	RequiredSkills []SkillAPI `json:"requiredSkills,omitempty" validate:"required,dive,required"`
}

type Status int

const (
	StatusConfirmed Status = iota + 1
	StatusAccepted
	StatusInProgress
	StatusArchived
	StatusCanceled
)

func (s Status) String() string {
	statuses := []string{"confirmed", "accepted", "in_progress", "archived", "canceled"}
	if s < StatusConfirmed || s > StatusCanceled {
		return "undefined"
	}
	return statuses[s]
}
