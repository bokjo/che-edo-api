package model

import (
	"database/sql"
	"log"
	"time"

	//Postgres driver
	_ "github.com/lib/pq"
)

// DBConnect creates new DB connection handler
func DBConnect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	// Retry logic
	//connected := false
	retries := 5

	for i := 1; i <= retries; i++ {

		if err = db.Ping(); err == nil {
			//return nil, err
			//connected = true
			log.Print("[DB] Connection established!")
			break
		} else {
			log.Printf("[DB] Connection retry: %d", i)
			time.Sleep(time.Second * 5)
		}

		// TEMP - implement better logic! and make custom error message!
		if i == 5 {
			return nil, err
		}
	}

	return db, nil
}
