package category

import "github.com/joshsoftware/golang-boilerplate/db"

type updateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type createRequest struct {
	Name string `json:"name"`
}

type findByIDResponse struct {
	Category db.Category `json:"category"`
}

type listResponse struct {
	Categories []db.Category `json:"categories"`
}

func (cr createRequest) Validate() (err error) {
	if cr.Name == "" {
		return errEmptyName
	}
	return
}

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}
