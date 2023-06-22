package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/porky256/mock-interview-tbr/internal/models/apimodels"
)

func RegisterValidator() (*validator.Validate, error) {
	x := validator.New()
	err := x.RegisterValidation("status_custom_validation", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(apimodels.Status)
		return value.String() != "unknown"
	})
	if err != nil {
		return nil, fmt.Errorf("can't register validator: %w", err)
	}

	err = x.RegisterValidation("user_status_custom_validation", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(apimodels.UserStatus)
		return value.String() != "unknown"
	})
	if err != nil {
		return nil, fmt.Errorf("can't register validator: %w", err)
	}

	return x, nil
}
