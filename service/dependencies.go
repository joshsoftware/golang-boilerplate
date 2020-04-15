package service

import "joshsoftware/golang-boilerplate/db"

type Dependencies struct {
	Store db.Storer
	// define other service dependencies
}
