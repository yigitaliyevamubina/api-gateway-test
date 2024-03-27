package storage

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type UserRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Product struct {
	Id          string   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Amount      int32   `json:"amount"`
}

// User info validation
func (u *UserRequest) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
		validation.Field(&u.LastName, validation.Required, validation.Length(5, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
	)
}

type Message struct {
	Message string `json:"message"`
}
