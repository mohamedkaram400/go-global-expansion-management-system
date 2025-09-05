package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/config"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	// "github.com/mohamedkaram400/go-global-expansion-management-system/db/seeders"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/routes"
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

	// 4. Connect to Redis
	if err := conn.ConnectRedis(config.RedisHost); err != nil {
		log.Fatal("❌ Failed to connect Redis:", err)
	}

	defer mongo.Disconnect(context.Background())

	// seeders.SeedAdminUser(mysql)  // Run only once

	// 5. Service, Repo and Handlers
	// Auth Module
	authRepo := repositories.NewAuthRepo(mysql)
	authService := services.NewAuthService(authRepo)
	authHandler := http.NewAuthHandler(authService)

	// Client Module
	clientRepo := repositories.NewClientRepo(mysql)
	clientService := services.NewClientService(clientRepo)
	clientHandler := http.NewClientHandler(clientService)

	// User Module
	userRepo := repositories.NewUserRepo(mysql)
	userService := services.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	// Vendor Module
	vendorRepo := repositories.NewVendorRepo(mysql)
	vendorService := services.NewVendorService(vendorRepo)
	vendorHandler := http.NewVendorHandler(vendorService)

	// 6. Init router
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery())

	// 7. Versioned API group
	v1 := router.Group("/api/v1")

	// 8. Register routes by module
	routes.RegisterAuthRoutes(v1,   authHandler)
	routes.RegisterClientRoutes(v1, clientHandler)
	routes.RegisterVendorRoutes(v1, vendorHandler)
	routes.RegisterUserRoutes(v1, userHandler)

	// 9. Test server
	TestServer(router)

	// 10. Start server
	startServer(router, config)
}

func TestServer(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}

func startServer(router *gin.Engine, config *config.Config) {
	if err := router.Run(config.Port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
	log.Println("🚀 App started on port", config.Port)
}
