package main

import (
	"log"
  "context"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/config"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
)


func main() {

	// 1. Load Config
	config := config.LoadConfig()

  // 2. Connect to MongoDB
	mongo, err := conn.ConnectMongo(config.MongoURI)	
	if err != nil {
		log.Fatal("❌ Failed to connect Mongo:", err)
	}

  // 3. Connect to MySQL
  mysql, err := conn.ConnectMySQL(config.MySQLURI)
  if err != nil {
		log.Fatal("❌ Failed to connect MySQL:", err)
  }

  defer mongo.Disconnect(context.Background())

	
  // 4. Start server
  startServer(mysql, config)
}

func startServer(mysql *gorm.DB, config *config.Config) {
  // Init routes
  router := gin.Default()
  router.SetTrustedProxies(nil)
  router.Use(gin.Logger(), gin.Recovery())


  // Just a test route
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
  })

  router.Run(config.Port)
  log.Println("🚀 App started on port", config.Port)
}