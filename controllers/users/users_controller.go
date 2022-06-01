package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/idoyudha/bookstore_oauth-go/oauth"
	"github.com/idoyudha/bookstore_utils-go/rest_errors"
)

func getUserId(userIdParam string) (int64, *rest_errors.RestErr) {
	// take userid from url
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
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
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// function below is handle same from line 16 to 27
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	// Handle JSON error
	// 	return
	// }

	// 2. process data
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		// handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	// 3. return result
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

func Update(c *gin.Context) {
	// take userid from url
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	// take json body
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
