package conn

import (
	"log"

	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 🔇 turn off all SQL logs
	})
	if err != nil {
		return nil, err
	}
	
	log.Println("✅ Connected to MySQL successfully")
	return db, nil
}