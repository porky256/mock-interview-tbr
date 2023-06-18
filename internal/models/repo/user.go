package repomodel

import "time"

type User struct {
	ID          int
	Username    string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Phone       string
	UserStatus  int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserSkill struct {
	ID        int
	SkillID   int
	UserID    int
	Score     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
