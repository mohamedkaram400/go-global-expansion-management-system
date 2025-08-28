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

  // Service, Repo and Handlers
	// 2. Init Repo
	authRepo := repositories.NewAuthRepo(mysql)

	// 3. Init Service
	authService := services.NewAuthService(authRepo)

	// 4. Init Handler
	authHandler := http.NewAuthHandler(authService)


  // Routes
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery())

	// Register routes
	TestServer(router)
	// RegisterRoutes(router)

	// Start server
	startServer(router, config)
}

func TestServer(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}

// func RegisterRoutes(router *gin.Engine) {
// 	authHandler := NewAuthHandler()
// 	router.POST("/register", authHandler.Register)
// 	router.POST("/login", authHandler.Login)
// }

func startServer(router *gin.Engine, config *config.Config) {
	if err := router.Run(config.Port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
	log.Println("🚀 App started on port", config.Port)
}