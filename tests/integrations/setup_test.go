package integrations

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func TestMain(m *testing.M) {
	// ✅ Go one step back from the current working directory
	wd, _ := os.Getwd()
	projectRoot := filepath.Join(wd, "..", "..") // Move 2 levels up: from tests/integrations → project root

	// ✅ Create the tmp directory at the project root
	dbDir := filepath.Join(projectRoot, "tmp")
	dbPath := filepath.Join(dbDir, "my_test_database.db")

	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			log.Fatalf("❌ failed to create tmp directory: %v", err)
		}
	}

	var err error
	TestDB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect to test DB: %v", err)
	}

	models := []interface{}{
		&entities.User{},
		&entities.Client{},
		&entities.Vendor{},
		&entities.VendorStatus{},
		&entities.Project{},
		&entities.Match{},
	}

	if err := TestDB.AutoMigrate(models...); err != nil {
		log.Fatalf("❌ failed to migrate test DB: %v", err)
	}

	log.Println("✅ Test DB ready at:", dbPath)

	code := m.Run()
	os.Exit(code)
}
