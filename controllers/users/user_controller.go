package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/balwaninitu/nitu-s_bookstore_api/domain/users"
	"github.com/balwaninitu/nitu-s_bookstore_api/services"
	"github.com/balwaninitu/nitu-s_bookstore_api/utils/errors"
	"github.com/gin-gonic/gin"
)

//controller package will be first or intenal layer of mvc pattern
//all rquest get handle by controllers through end points
//entry point of every request is controller
/*only at controller we define http server, its convenience to keep http
at only one point so in future if need to change only one package will be affected
here only application and controller package using http server*/

// bytes, err := ioutil.ReadAll(c.Request.Body)
// if err != nil {
// 	//TODO:Handle error
// 	return
// }
// if err := json.Unmarshal(bytes, &user); err != nil {
// 	fmt.Println(err.Error())
// 	//TODO:Handle json error
// 	return
// }
//other way of getting json data alternate to unmarshal func
//above can be commented out, shouldbind json doing same underhood as above

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, useErr := strconv.ParseInt(userIdParam, 10, 64)
	if useErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")

	}
	return userId, nil
}
func Create(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//fmt.Println(err)
		return
	}
	//saving in database
	//controller is no incharge of databse, it all take care by services
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(user)
	//fmt.Println(err)
	//fmt.Println(string(bytes))
	//c.String(http.StatusNotImplemented, "Implement me!")
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	//c.String(http.StatusNotImplemented, "Implement me!")
	//if true then its public request otherwise private request
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))

}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
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
	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

	/*1st approach, if we have single end point that takes slice of user then
	can use below code but problem arise when you have more endpoint returning a list
	of users then you need to copy below code everytime for all endpoint so its better
	to turn it into function and call everytime */

	// result := make([]interface{}, len(users))
	// for index, user := range users{
	// 	result[index] = user.Marshall(c.GetHeader("X-Public") == "true")
	// }

}
