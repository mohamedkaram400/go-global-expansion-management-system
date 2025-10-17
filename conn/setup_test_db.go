package conn

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var TestDB *gorm.DB

func SetupTestDB() *gorm.DB {
	// ✅ Get current working directory
	wd, _ := os.Getwd()

	// ✅ Move up one or two levels as needed
	projectRoot := filepath.Join(wd, "..")
	dbDir := filepath.Join(projectRoot, "tmp")
	dbPath := filepath.Join(dbDir, "my_test_database.db")

	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			log.Fatalf("❌ failed to create tmp directory: %v", err)
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("❌ failed to migrate test DB: %v", err)
	}

	// log.Println("✅ Test DB ready at:", dbPath)

	TestDB = db // Save it globally if needed later
	return db
}
