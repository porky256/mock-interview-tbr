package models

import "time"

type UserMatchRepo struct {
	ID         int
	UserAsker  int
	UserMatch  int
	MatchScore int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type MatchRequestRepo struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillInRequestRepo struct {
	ID        int
	RequestID int
	SkillID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InterviewRepo struct {
	ID            int
	MatchID       int
	Status        int
	InterviewDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
