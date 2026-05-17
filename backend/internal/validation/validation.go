package validation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationService interface {
	ValidateJSON(any) error
}

type validationService struct {
	validator *validator.Validate
}

func NewService() ValidationService {
	return validationService{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s validationService) ValidateJSON(data any) error {
	if err := s.validator.Struct(data); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = fe.Field() + " field " + msgForTag(fe)
			}
			return errors.New(strings.Join(out, ", "))
		}
		return err
	}
	return nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "required"
	case "min":
		return "must have the length of at least " + fe.Param()
	default:
		return fe.Tag() + " - unknown"
	}
}
