package seeders

import (
	"errors"
	"log"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/pkg"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB) {
	// check if admin already exists
	var admin entities.User
	result := db.Where("role = ?", "Admin").Take(&admin)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		hashedPassword, _ := pkg.HashPassword("Admin@123")

		admin = entities.User{
			Name:     "Admin",
			Email:    "mohamed@admin.com",
			Password: hashedPassword,
			Role:     "Admin",
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Fatalf("❌ Failed to seed admin user: %v", err)
		}

		log.Println("✅ Admin user seeded successfully")
		return
	}

	if result.Error != nil {
		log.Fatalf("❌ Failed to check admin user: %v", result.Error)
	}

	log.Println("⚠️ Admin user already exists, skipping seeding")
}
