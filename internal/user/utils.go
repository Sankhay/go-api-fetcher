package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func generateCreateUserValidationErrors(err error) string {
	var validationErrors string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			var customMsg string
			switch e.Tag() {
			case "required":
				customMsg = fmt.Sprintf("The field '%s' is required.", e.Field())
			case "email":
				customMsg = fmt.Sprintf("The field '%s' must be a valid email address.", e.Field())
			default:
				customMsg = fmt.Sprintf("The field '%s' failed validation for tag: %s", e.Field(), e.Tag())
			}
			validationErrors += customMsg + " "
		}
	}
	return validationErrors
}
