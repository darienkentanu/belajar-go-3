package repositories

import (
	"belajar-go-3/mvc-with-db/config"
	"belajar-go-3/mvc-with-db/libraries"
	"database/sql"
)

var db *sql.DB
var err error

func init() {
	// Create an sql.DB and check for errors
	db, err = sql.Open(config.GetString("database.driverName"), config.GetString("database.dataSourceName"))
	libraries.CheckError(err)

	// Test the connection to the database
	err = db.Ping()
	libraries.CheckError(err)
}
