package users

import (
	"fmt"

	"github.com/balwaninitu/nitu-s_bookstore_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

/*dao is data acceess object, as entire logic to peresist and
retrieve user from given database i.e access layer to database
It is the only point in entire application where we are working
interacting with  database
 if in future there need to change database then only change will
 be in dao file in whole application */

/*methods are better to assign directly to user struct
instead of function */

//get user from database based on user id
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

//save user in database
func (user *User) Save() *errors.RestErr {
	//checking if user id already existing
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))

	}
	usersDB[user.Id] = user
	return nil
}
