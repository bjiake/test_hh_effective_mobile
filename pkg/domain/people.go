package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

// People represents a people entity
type People struct {
	// ID is the unique identifier of the people
	ID int64 `json:"id" validate:"required" swagger:"description:Unique identifier of the people"`

	// Name is the name of the people
	Name string `json:"name" validate:"required,latin-cyrillic" swagger:"description:Name of the people"`

	// SurName is the surname of the people
	SurName string `json:"surName" validate:"required,latin-cyrillic" swagger:"description:Surname of the people"`

	// Patronymic is the patronymic of the people
	Patronymic string `json:"patronymic" validate:"required,latin-cyrillic" swagger:"description:Patronymic of the people"`
}

func (p People) Validate() error {
	validate := validator.New()

	// Register the custom validation function (only if not already registered)
	if err := validate.RegisterValidation("latin-cyrillic", validateLatinCyrillic); err != nil {
		return err
	}

	err := validate.Struct(p)
	if err != nil {
		// Handle validation errors
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return fmt.Errorf("people validation errors: %s", strings.Join(validationErrors, ", "))
	}

	return nil // No validation errors
}

// Custom validation function for Latin and Cyrillic characters
func validateLatinCyrillic(fl validator.FieldLevel) bool {
	// Regular expression to match Latin and Cyrillic characters
	regex := regexp.MustCompile(`^[\p{Latin}\p{Cyrillic}\s]+$`)
	return regex.MatchString(fl.Field().String())
}

type UpdatePeople struct {
	ID         int64  `json:"id" validate:"required"`
	Name       string `json:"name" validate:"omitempty,latin-cyrillic"`
	SurName    string `json:"surName" validate:"omitempty,latin-cyrillic"`
	Patronymic string `json:"patronymic" validate:"omitempty,latin-cyrillic"`
}

func (p UpdatePeople) Validate() error {
	validate := validator.New()

	// Register the custom validation function (only if not already registered)
	if err := validate.RegisterValidation("latin-cyrillic", validateLatinCyrillic); err != nil {
		return err
	}

	err := validate.Struct(p)
	if err != nil {
		// Handle validation errors
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return fmt.Errorf("UpdatePeople validation errors: %s", strings.Join(validationErrors, ", "))
	}

	return nil // No validation errors
}
