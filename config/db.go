package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	dbInfo := "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable"
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
		}).Fatal("failed to ping to database")
	}
	db = conn
	err = createUsersTable(db)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("failed to create users table")
	}
}

func createUsersTable(db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		email VARCHAR(300) NOT NULL,
		password VARCHAR(300) NOT NULL
	);`
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}
