package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authHandler *http.AuthHandler) {
    auth := rg.Group("/auth")
    v2 := auth.Group("/v2")

	v2.Use(middlewares.JWTAuth()) 

    {
        auth.POST("/register", authHandler.Register)
        auth.POST("/login", authHandler.Login)
        v2.POST("/logout", authHandler.Logout)
    }
}
