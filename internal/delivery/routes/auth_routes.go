package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authHandler *http.AuthHandler) {
	auth := rg.Group("/auth")

	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
        auth.POST("/logout", middlewares.JWTAuth(), authHandler.Logout)
	}
}
