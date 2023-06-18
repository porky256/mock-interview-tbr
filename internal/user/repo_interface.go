package user

import (
	repomodel "github.com/porky256/mock-interview-tbr/internal/models"
)

// DatabaseUserProvider provider for db related functions for users
type DatabaseUserProvider interface {
	InsertUser(u repomodel.UserRepo) error
	GetUserByID(id int) (*repomodel.UserRepo, error)
}
