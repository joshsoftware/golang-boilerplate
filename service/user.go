package service

import (
	"errors"
	"joshsoftware/golang-boilerplate/db"
)

func ValidateUserAge(u db.User) (err error) {
	if u.Age <= 0 {
		err = errors.New("age cannot be less than equal to 0")
		return
	}
	return
}
