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
func CreateUser(c *gin.Context) {
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
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(user)
	//fmt.Println(err)
	//fmt.Println(string(bytes))
	//c.String(http.StatusNotImplemented, "Implement me!")
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, useErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if useErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	//c.String(http.StatusNotImplemented, "Implement me!")
	c.JSON(http.StatusOK, user)

}
