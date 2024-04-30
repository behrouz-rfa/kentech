// Package validate contains the support for validating models.
package validate

import (
	"errors"
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

// translator is a cache of locale and translation information.
var translator ut.Translator

func Init() {
	// Instantiate a validator.
	validate = validator.New()
}

// Check validates the provided model against it's declared tags.
func Check(val any) error {
	if err := validate.Struct(val); err != nil {
		// Use a type assertion to get the real error value.
		var verrors validator.ValidationErrors
		if !errors.As(err, &verrors) {
			return err
		}

		fields := make(FieldErrors, 0, len(verrors))

		for _, verror := range verrors {
			field := FieldError{
				Field: verror.Field(),
				Error: verror.Translate(translator),
			}
			fields = append(fields, field)
		}

		return fields
	}

	return nil
}

// CheckEmail validates that the string is an email.
func CheckEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidateFaxPhone validates that the string is an fax.
func ValidateFaxPhone(fl string) bool {
	// Fax numbers can be in different formats, but generally contain numbers and possibly dashes or parentheses.
	// Here we use a regular expression to validate a common format for North American fax numbers: (123) 456-7890.
	faxRegex := regexp.MustCompile(`^\+?[0-9]\d{1,14}$`)
	return faxRegex.MatchString(fl)
}
