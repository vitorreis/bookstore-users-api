package services

import (
	"github.com/vitorreis/bookstore-users-api/domain/users"
	"github.com/vitorreis/bookstore-users-api/util/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}

func GetUser() {

}

func FindUser() {

}
