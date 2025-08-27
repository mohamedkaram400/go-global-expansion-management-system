package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/config"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
)


func main() {

	// 1. Load Config
	config := config.LoadConfig()

  // 2. Connect to MySQL
  mysql, err := conn.ConnectMySQL(config.MySQLURI)
  if err != nil {
		log.Fatal("❌ Failed to connect MySQL:", err)
  }

  // 3. Connect to MongoDB
	mongo, err := conn.ConnectMongo(config.MongoURI)	
	if err != nil {
		log.Fatal("❌ Failed to connect Mongo:", err)
	}

  defer mongo.Disconnect(context.Background())

  // migrate (if needed)
	// conn.AutoMigrate(&entities.Client{})

  // Service
	// 2. Init Repo
	authRepo := repositories.NewAuthRepo(mysql)

	// 3. Init Service
	authService := services.NewAuthService(authRepo)

	// 4. Init Handler
	authHandler := http.NewAuthHandler(authService)

	router := gin.Default()
  router.SetTrustedProxies(nil)
  router.Use(gin.Logger(), gin.Recovery())


  // Just a test route
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
  })

  // Register routes
  router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

  router.Run(config.Port)
  if err := router.Run(config.Port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
  log.Println("🚀 App started on port", config.Port)

  // 4. Start server
  // startServer(mysql, config)
}

// func startServer(mysql *gorm.DB, config *config.Config) {
  //   // Init routes
  //   router := gin.Default()
  //   router.SetTrustedProxies(nil)
  //   router.Use(gin.Logger(), gin.Recovery())

  //   // Just a test route
  //   router.GET("/ping", func(c *gin.Context) {
  //     c.JSON(200, gin.H{"message": "pong"})
  //   })

  //   // Register routes
  //   router.POST("/register", authHandler.Register)
  // 	router.POST("/login", authHandler.Login)

  //   router.Run(config.Port)
  //   log.Println("🚀 App started on port", config.Port)
// }