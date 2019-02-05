package db

import "errors"

var (
	ErrCategoryNotExist = errors.New("Category does not exist in db")
)
