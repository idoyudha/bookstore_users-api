// data transfer object = transfering from the database and the code

package users

import (
	"strings"

	"github.com/idoyudha/bookstore_utils-go/rest_errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

// this is a Validate() FUNCTION (*not using this)
func Validate(user *User) *rest_errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Invalid email address")
	}
	return nil
}

// this is a Validate() METHOD attached to User struct
func (user *User) Validate() *rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("Invalid password")
	}
	return nil
}
