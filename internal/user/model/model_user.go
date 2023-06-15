package model

import "github.com/porky256/mock-interview-tbr/internal/skill/model"

type User struct {
	Id          int64        `json:"id,omitempty"`
	Username    string       `json:"username,omitempty"`
	FirstName   string       `json:"firstName,omitempty"`
	LastName    string       `json:"lastName,omitempty"`
	Email       string       `json:"email,omitempty"`
	Password    string       `json:"password,omitempty"`
	Phone       string       `json:"phone,omitempty"`
	UserStatus  int32        `json:"userStatus,omitempty"`
	Description string       `json:"description,omitempty"`
	UsersSkills []UsersSkill `json:"usersSkills,omitempty"`
}

type UsersSkill struct {
	Id    int64        `json:"id,omitempty"`
	Skill *model.Skill `json:"skill,omitempty"`
	Score int32        `json:"score,omitempty"`
}
