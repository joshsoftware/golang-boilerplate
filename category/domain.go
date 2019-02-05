package category

import "github.com/joshsoftware/golang-boilerplate/db"

type category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type findByIDResponse struct {
	Category db.Category `json:"category"`
}

type listResponse struct {
	Categories []db.Category `json:"categories"`
}

func (c category) Validate() (err error) {
	if c.Name == "" {
		return errEmptyName
	}

	return
}
