package category

import "errors"

var (
	errEmptyID      = errors.New("Category ID must be present")
	errEmptyName    = errors.New("Category name must be present")
	errNoCategories = errors.New("No categories present")
	errNoCategoryId = errors.New("Category is not present")
)
