package conn

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewTestDB creates and returns a SQLite in-memory DB for testing
func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("tmp/my_test_database.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to SQLite: %v", err)
	}

	log.Println("✅ Connected to SQLite successfully")
	return db, nil
}