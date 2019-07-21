package handlers

import (
	"reflect"
	"testing"
)

type StructTest struct {
	Name          string `validate:"required"`
	Email         string `validate:"required,email"`
	Age           int    `validate:"gte=18"`
	Document      string `validate:"document"`
	BirthdayMonth int    `validate:"min=1,max=12"`
}

func TestValidate(t *testing.T) {
	t.Run("returns a error if a required field is empty", func(t *testing.T) {
		s := StructTest{Name: "Jessica", Age: 21, Document: "98741303269", BirthdayMonth: 8}
		got := Validate(s)
		want := []string{"Email: is required"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns a error if a email is not valid", func(t *testing.T) {
		s := StructTest{Name: "Jessica", Age: 21, Email: "jessicapaz", Document: "98741303269", BirthdayMonth: 8}
		got := Validate(s)
		want := []string{"Email: must be a valid email"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns a error if a field is not greater than or equal to something", func(t *testing.T) {
		s := StructTest{Name: "Jessica", Age: 10, Email: "j@mail.com", Document: "98741303269", BirthdayMonth: 8}
		got := Validate(s)
		want := []string{"Age: must be greater than or equal to 18"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns a error if a field is a document", func(t *testing.T) {
		s := StructTest{Name: "Jessica", Age: 19, Email: "j@mail.com", Document: "083901239012878", BirthdayMonth: 8}
		got := Validate(s)
		want := []string{"Document: must be equal to 14"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns a error if a field is less than a minimum", func(t *testing.T) {
		s := StructTest{Name: "Jessica", Age: 19, Email: "j@mail.com", Document: "98741587390", BirthdayMonth: 0}
		got := Validate(s)
		want := []string{"BirthdayMonth: must be greater than or equal to 1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns a error if a field is greater than a maximum", func(t *testing.T) {
		s := StructTest{Name: "Jessica", Age: 19, Email: "j@mail.com", Document: "98741587309", BirthdayMonth: 16}
		got := Validate(s)
		want := []string{"BirthdayMonth: must be less than or equal to 12"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestValidateSortQueryParam(t *testing.T) {
	got := ValidateSortQueryParam("document banana")
	want := "Not a valid option"

	if got.Error() != want {
		t.Errorf("got %s want %s", got, want)
	}
}
