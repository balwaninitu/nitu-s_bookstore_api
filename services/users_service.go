package services

import (
	"github.com/balwaninitu/nitu-s_bookstore_api/domain/users"
	"github.com/balwaninitu/nitu-s_bookstore_api/utils/errors"
)

//entire business logic of application will be on services
//for validation user should have valid email id
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
