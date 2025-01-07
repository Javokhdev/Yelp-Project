package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"yalp/internal/config"
)

func Connect(cnf *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.DB_HOST, cnf.DB_PORT, cnf.DB_USER, cnf.DB_PASSWORD, cnf.DB_NAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database not reachable: %v", err)
	}

	return db, nil
}
