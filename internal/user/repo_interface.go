package user

import repomodel "github.com/porky256/mock-interview-tbr/internal/models/repo"

// DatabaseUserProvider provider for db related functions for users
type DatabaseUserProvider interface {
	InsertUser(u repomodel.User) error
	GetUserByID(id int) (*repomodel.User, error)
}
