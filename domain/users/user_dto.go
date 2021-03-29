package users

import (
	"strings"

	"github.com/balwaninitu/nitu-s_bookstore_api/utils/errors"
)

//core of application
/*dto stands for data transfer object, here transferring data from persistent
layer(database) to application */
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//for validation user should have valid email id
// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}
// 	return nil
// }
//method is better than above func
//assigning method validate to user struct
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

//
