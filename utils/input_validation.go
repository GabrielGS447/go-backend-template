package utils

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateInput(input interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	return validate.Struct(input)
}

func validateDate(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	// this regex is not perfect, but it's good enough for this example
	regex, _ := regexp.Compile(`^\d{4}-\d{2}-\d{2}$`)
	return regex.MatchString(date)
}

func GetValidationErrors(err error) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, e := range ve {
			out[i] = getErrorMsg(e)
		}
		return out
	}
	return nil
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	case "date":
		return fe.Field() + " must be a valid date in the format YYYY-MM-DD"
	default:
		return "An unknown error occurred"
	}
}
