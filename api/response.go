package api

import (
	"encoding/json"
	"net/http"

	"github.com/joshsoftware/golang-boilerplate/app"
)

type Response struct {
	Message string `json:"message"`
}

func Error(rw http.ResponseWriter, status int, response interface{}) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		app.GetLogger().Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func Success(rw http.ResponseWriter, status int, response interface{}) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		app.GetLogger().Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
