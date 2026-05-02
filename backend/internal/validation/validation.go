package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationService struct {
	validator *validator.Validate
}

func NewService() ValidationService {
	return ValidationService{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s ValidationService) ValidateJSON(data any) error {
	if err := s.validator.Struct(data); err != nil {
		return fmt.Errorf("failed to validate struct: %w", err)
	}
	return nil
}
