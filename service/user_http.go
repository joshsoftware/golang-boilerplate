package service

import (
	"joshsoftware/golang-boilerplate/api"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

// @Title listUsers
// @Description list all User
// @Router /users [get]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func listUsersHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		users, err := deps.Store.ListUsers(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			api.Error(http.StatusInternalServerError, api.Response{Message: "Error fetching data"}, rw)
			return
		}

		api.Success(http.StatusOK, users, rw)
	})
}
