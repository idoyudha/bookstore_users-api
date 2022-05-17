// data access object = logic that we need to execute
// if we have an error about database query, we just look into dao file

package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	"fmt"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser) // validate initial query
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // close connection if we don't have any error

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser) // validate initial query
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // close connection

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
	}

	// result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated) same with above and below

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId

	return nil
}
