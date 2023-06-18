package dal

import (
	dalModel "github.com/porky256/mock-interview-tbr/internal/dal/models"
)

// UserDatabaseProvider provider for db related functions for users
type UserDatabaseProvider interface {
	InsertUser(u dalModel.User) error
	GetUserByID(id int) (*dalModel.User, error)
}

// MatchDatabaseProvider provider for db related functions for matches
type MatchDatabaseProvider interface {
	GetUserMatchByID(id int) (*dalModel.UserMatch, error)
}

// SkillDatabaseProvider provider for db related functions for skills
type SkillDatabaseProvider interface {
	GetSkillByID(id int) (*dalModel.Skill, error)
}

// GlobalDatabaseProvider provider for all db related functions
type GlobalDatabaseProvider interface {
	UserDatabaseProvider
	MatchDatabaseProvider
	SkillDatabaseProvider
}
