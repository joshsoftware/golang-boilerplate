package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joshsoftware/golang-boilerplate/config"
	_ "github.com/lib/pq"
)

var ErrFindingDriver = errors.New("no migrate driver instance found")

func RunMigrations() error {
	dbConfig := config.Database()

	db, err := sql.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		return err
	}

	driver, err := getDBDriverInstance(db, dbConfig.Driver())
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(GetMigrationPath(), dbConfig.Driver(), driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func getDBDriverInstance(db *sql.DB, driver string) (database.Driver, error) {
	switch driver {
	case "postgres":
		return postgres.WithInstance(db, &postgres.Config{})
	default:
		return nil, ErrFindingDriver
	}
}

func CreateMigrationFile(filename string) error {
	if len(filename) == 0 {
		return errors.New("filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", config.MigrationPath(), timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", config.MigrationPath(), timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}

	fmt.Printf("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	fmt.Printf("created %s\n", downMigrationFilePath)
	return nil
}

func RollbackMigrations(s string) error {
	steps, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	m, err := migrate.New(GetMigrationPath(), config.Database().ConnectionURL())
	if err != nil {
		return err
	}

	err = m.Steps(-1 * steps)
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = f.Close()

	return err
}

func GetMigrationPath() string {
	return fmt.Sprintf("file://%s", config.MigrationPath())
}
