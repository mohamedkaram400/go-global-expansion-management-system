package conn

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 🔇 turn off all SQL logs
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Mysql: %v", err)
	}

	log.Println("✅ Connected to MySQL successfully")
	return db, nil
}