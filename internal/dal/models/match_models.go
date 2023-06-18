package dalmodel

import "time"

type UserMatch struct {
	ID         int
	UserAsker  int
	UserMatch  int
	MatchScore int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type MatchRequest struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillInRequest struct {
	ID        int
	RequestID int
	SkillID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Interview struct {
	ID            int
	MatchID       int
	Status        int
	InterviewDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
