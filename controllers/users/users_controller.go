package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// 1. take request body
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(user)
	if err != nil {
		//TODO: Handle error
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		// Handle JSON error bad request
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// function below is handle same from line 16 to 27
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	// Handle JSON error
	// 	return
	// }

	// 2. process data
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	// 3. return result
	c.JSON(http.StatusCreated, result)
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not yet implemented!")
}
