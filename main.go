package main

// @APITitle Main
// @APIDescription Main API for Microservices in Go!

import (
	"fmt"
	"strconv"

	"github.com/urfave/negroni"

	logger "github.com/sirupsen/logrus"
	"joshsoftware/golang-boilerplate/config"
	"joshsoftware/golang-boilerplate/db"
	"joshsoftware/golang-boilerplate/service"
)

func main() {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	config.Load()

	store, err := db.Init()
	if err != nil {
		logger.Error("Database init failed", err)
		return
	}

	deps := service.Dependencies{
		Store: store,
	}

	// mux router
	router := service.InitRouter(deps)

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
}
