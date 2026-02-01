package database

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id VARCHAR(36) NOT NULL,
		price DECIMAL(10,2) NOT NULL,
		tax DECIMAL(10,2) NOT NULL,
		final_price DECIMAL(10,2) NOT NULL,
		PRIMARY KEY (id)
	);
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
