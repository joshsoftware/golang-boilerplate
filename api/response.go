package api

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type Response struct {
	Message string `json:"message"`
}

func Error(status int, response interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func Success(status int, response interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
