package user

import (
	repomodel "github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// DatabaseUserProvider provider for db related functions for users
type DatabaseUserProvider interface {
	InsertUser(user repomodel.UserRepo) (int, error)
	GetUserByID(id int) (*repomodel.UserRepo, error)
	GetUserByUsername(username string) (*repomodel.UserRepo, error)
	UpdateUser(user repomodel.UserRepo) error
	DeleteUserByID(id int) error

	InsertUserSkill(skill repomodel.UserSkillRepo) (int, error)
	GetUserSkillByID(id int) (*repomodel.UserSkillRepo, error)
	GetUsersSkillsByUserID(userID int) ([]repomodel.UserSkillRepo, error)
	UpdateUserSkill(skill repomodel.UserSkillRepo) error
	DeleteUserSkillByID(id int) error
	DeleteUserSkillByUserID(userID int) error
}
