package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"joshsoftware/golang-boilerplate/config"
)

// Create your schema here (sample provided below)
// If this schema is too big, put it in a schema.go file
var schema = `
CREATE TABLE IF NOT EXISTS users (
	name text,
	age integer
);`

type pgStore struct {
	db *sqlx.DB
}

func Init() (s Storer, err error) {
	uri := config.ReadEnvString("DB_URI")

	conn, err := sqlx.Connect("postgres", uri)
	if err != nil {
		logger.Error("Cannot initialize database", err)
		return
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	conn.MustExec(schema)

	return &pgStore{conn}, nil
}
