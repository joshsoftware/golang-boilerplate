package main

// @APITitle Main
// @APIDescription Main API for Microservices in Go!

import (
	"context"
	"fmt"
	"joshsoftware/golang-boilerplate/config"
	"joshsoftware/golang-boilerplate/db"
	"joshsoftware/golang-boilerplate/logger"
	"joshsoftware/golang-boilerplate/service"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

func main() {
	config.Load()

	cliApp := cli.NewApp()
	cliApp.Name = config.AppName()
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				return startApp()
			},
		},
		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				return db.RunMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp() (err error) {
	_, err = logger.SetupLogger(config.Env())
	if err != nil {
		return
	}

	store, err := db.Init()
	if err != nil {
		logger.Errorw(context.Background(), "Database init failed", "err", err.Error())
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

	port := config.AppPort() // This can be changed to the service port number via environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
	return
}
