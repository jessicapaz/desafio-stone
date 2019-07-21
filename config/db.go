package config

import (
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	dbInfo := os.Getenv("DATABASE_URL")
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
	err = createInvoicesTable(db)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("failed to create invoices table")
	}
}

func createUsersTable(db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		email VARCHAR(300) NOT NULL UNIQUE,
		password VARCHAR(300) NOT NULL
	);`
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func createInvoicesTable(db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS invoices(
		id SERIAL PRIMARY KEY,
		reference_month INTEGER NOT NULL,
		reference_year INTEGER NOT NULL,
		document VARCHAR(14) NOT NULL,
		description VARCHAR(256) NOT NULL,
		amount NUMERIC(16, 2) NOT NULL,
		is_active SMALLINT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		deactivated_at TIMESTAMP
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
