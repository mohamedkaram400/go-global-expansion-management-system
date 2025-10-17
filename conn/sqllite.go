package conn

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewTestDB creates and returns a SQLite in-memory DB for testing
func ConnectSQLite() (*gorm.DB, error) {
	wd, _ := os.Getwd()
    dbPath := filepath.Join(wd, "tmp", "my_test_database.db")

	log.Printf("Using SQLite file at: %s", dbPath)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to SQLite: %v", err)
	}

	log.Println("✅ Connected to SQLite successfully")
	return db, nil
}
