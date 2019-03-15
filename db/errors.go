package db

import "errors"

const (
	DuplicateData = "23505"
)

var (
	ErrCategoryNotExist          = errors.New("Category does not exist in db")
	errCategoryDuplicateKeyValue = errors.New("Duplicate data inserted into database table")
)
