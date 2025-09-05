package seeders

import (
	"log"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/pkg"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB) {
	// check if admin already exists
	var admin entities.User
	result := db.Where("role = ?", "Admin").First(&admin)

	if result.Error == gorm.ErrRecordNotFound {
		// create admin user
		hashedPassword, _ := pkg.HashPassword("Admin@123") 
		admin = entities.User{
			Name:     "Admin",
			Email:    "mohamed@admin.com",
			Password: hashedPassword,
			Role:     "Admin",
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Fatalf("Failed to seed admin user: %v", err)
		}
		log.Println("✅ Admin user seeded successfully")
	} else {
		log.Println("⚠️ Admin user already exists, skipping seeding")
	}
}