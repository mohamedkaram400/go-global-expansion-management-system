package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
)

func AuthRoutes(rg *gin.RouterGroup, authHandler *http.AuthHandler) {
    auth := rg.Group("/auth")
	// auth.Use(AuthMiddleware()) // all v1 routes require auth

    {
        auth.POST("/register", authHandler.Register)
        auth.POST("/login", authHandler.Login)
    }
}
