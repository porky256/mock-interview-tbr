package userservice

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/porky256/mock-interview-tbr/internal/models/apimodels"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
	"github.com/porky256/mock-interview-tbr/internal/user"
)

type Service struct {
	DP       user.DatabaseUserProvider
	Validate *validator.Validate
}

func NewService(provider user.DatabaseUserProvider, validate *validator.Validate) Service {
	return Service{
		DP:       provider,
		Validate: validate,
	}
}

func (s *Service) AddUser(user apimodels.UserAPI) error {
	err := s.Validate.Struct(user)
	if err != nil {
		return fmt.Errorf("user isn't valid: %w", err)
	}

	repoUser := repomodels.UserRepo{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		Phone:       user.Phone,
		UserStatus:  int(user.UserStatus),
		Description: user.Description,
	}

	_, err = s.DP.InsertUser(repoUser)
	if err != nil {
		return fmt.Errorf("can't insert user in database: %w", err)
	}

	return nil
}
