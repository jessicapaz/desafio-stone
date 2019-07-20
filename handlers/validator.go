package handlers

import (
	"errors"
	"strings"

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

func ValidateSortQueryParam(s string) error {
	if s == "" {
		return errors.New("Empty")
	}
	sortList := strings.Split(s, ",")
	for _, v := range sortList {
		sort := strings.Split(v, " ")
		options := []string{"document", "reference_year", "reference_month"}
		if stringInSlice(sort[0], options) == false {
			return errors.New("Not a valid option")
		}
		options = []string{"desc", "asc"}
		if len(sort) == 2 {
			if stringInSlice(sort[1], options) == false {
				return errors.New("Not a valid option")
			}
		}
	}
	return nil
}

func stringInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == strings.TrimSpace(s) {
			return true
		}
	}
	return false
}
