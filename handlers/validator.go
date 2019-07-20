package handlers

import (
	"github.com/go-playground/validator"
)

// Validate validates a struct and returns a list of errors
func Validate(i interface{}) []string {
	validate := validator.New()
	validate.RegisterAlias("document", "len=11|len=14")
	e := []string{}
	err := validate.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch tag := err.Tag(); tag {
			case "required":
				e = append(e, err.Field()+": is required")
			case "min":
				e = append(e, err.Field()+": must be greater than or equal to "+err.Param())
			case "max":
				e = append(e, err.Field()+": must be less than or equal to "+err.Param())
			case "document":
				e = append(e, err.Field()+": must be equal to "+err.Param())
			}
		}
	}
	return e
}
