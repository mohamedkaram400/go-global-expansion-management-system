package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/config"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"

	"github.com/mohamedkaram400/go-global-expansion-management-system/db/seeders"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/notifier"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	authRepo "github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/scheduler"
	entities "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	authService "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
	authHandler "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/routes/v1"
	authRoute "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/routes/v1/auth"
)

func main() {

	// 1. Load Config
	config := config.LoadConfig()

	// 2. Connect to MySQL
	mysql, err := conn.ConnectMySQL(config.MySQLURI)
	sqlDB, _ := mysql.DB()
	if err != nil {
		log.Fatal("❌ Failed to connect MySQL:", err)
	}

	// 3. Connect to MongoDB
	mongo, err := conn.ConnectMongo(config.MongoURI)
	if err != nil {
		log.Fatal("❌ Failed to connect Mongo:", err)
	}
	researchDocumentsollection := mongo.Database(config.DBName).Collection(config.CollectionName)

	// 4. Connect to Redis
	if err := conn.ConnectRedis(config.RedisHost); err != nil {
		log.Fatal("❌ Failed to connect Redis:", err)
	}

	// Close the services connection by the of method
	defer mongo.Disconnect(context.Background())
	defer conn.RedisClient.Close()
	defer sqlDB.Close()


	// ✅ Run AutoMigrate to create tables if not exist
	mysql.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	if err := mysql.AutoMigrate(
		&entities.User{},
		&entities.Client{},
		&entities.Vendor{},
		&entities.VendorStatus{},
		&entities.Project{},
		&entities.Match{},
	); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	mysql.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	log.Println("✅ Database migrated successfully")

	seeders.SeedAdminUser(mysql)  // Run only once

	// 5. Service, Repo and Handlers
	// User Auth Module
	authUserRepo := authRepo.NewUserAuthRepo(mysql)
	authUserService := authService.NewUserAuthService(authUserRepo)
	authUserHandler := authHandler.NewUserAuthHandler(authUserService)

	// Client Auth Module
	authClientRepo := authRepo.NewClientAuthRepo(mysql)
	authClientService := authService.NewClientAuthService(authClientRepo)
	authClientHandler := authHandler.NewClientAuthHandler(authClientService)

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

	// Project Module
	projectRepo := repositories.NewProjectRepo(mysql)
	projectService := services.NewProjectService(projectRepo)
	projectHandler := http.NewProjectHandler(projectService)

	// Notifier Module
	notifier := notifier.NewSMTPNotifier(config.MailHost, config.MailPort, config.MailUser, config.MailPass, config.MailFrom)

	// Match Module
	matchRepo := repositories.NewMatchRepo(mysql)
	matchService := services.NewMatchService(matchRepo, projectService, notifier, clientService)
	matchHandler := http.NewMatchHandler(matchService)

	// Research Document Module with (Mongo DB)
	researchDocumentRepo := repositories.NewResearchDocumentRepo(researchDocumentsollection)
	researchDocumentService := services.NewResearchDocumentService(researchDocumentRepo)
	researchDocumentHandler := http.NewResearchDocumentHandler(researchDocumentService, projectService)

	// Analytics Module with (MySql, Mongo DB)
	analyticsService := services.NewAnalyticsService(matchRepo, researchDocumentRepo, projectRepo)
	AnalyticsHandler := http.NewAnalyticsHandler(analyticsService)

	// 6. Init router
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery())

	// 7. Versioned API group
	v1 := router.Group("/api/v1")

	// 8. Register routes by module
	authRoute.RegisterUserAuthRoutes(v1,   authUserHandler)
	authRoute.RegisterClientAuthRoutes(v1,   authClientHandler)
	routes.RegisterClientRoutes(v1, clientHandler)
	routes.RegisterVendorRoutes(v1, vendorHandler)
	routes.RegisterUserRoutes(v1, userHandler)
	routes.RegisterProjectRoutes(v1, projectHandler)
	routes.RegisterMatchRoutes(v1, matchHandler)
	routes.RegisterResearchDocumentRoutes(v1, researchDocumentHandler)
	routes.RegisterAnalyticsRoutes(v1, AnalyticsHandler)

	// 9. Add cron job for re-matching
	ctx := context.Background()
	jobManager := scheduler.NewJobManager(ctx)

	jobManager.RegisterJob(&scheduler.RefreshMatchesJob{MatchService: matchService, ProjectService: projectService})
	jobManager.RegisterJob(&scheduler.FlagExpiredSLAsJob{VendorService: vendorService})

	// jobManager.StartScheduler()

	// 10. Test server
	TestServer(router)

	// 11. Start server
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

