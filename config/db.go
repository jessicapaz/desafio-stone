package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

var db *sql.DB

func init() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbInfo := fmt.Sprintf("host=db user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)
	conn, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("failed to connect to postgres")
	}
	err = conn.Ping()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("unable to connect to database")
	}
	db = conn
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}
