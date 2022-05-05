// data access object = logic that we need to execute
// if we have an error about database query, we just look into dao file

package users

import (
	"bookstore_users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id] // get value from db by id
	if current != nil {         // if value is found, then action
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s alredy registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d alredy exist", user.Id))
	}
	usersDB[user.Id] = user // save user value to db
	return nil
}
