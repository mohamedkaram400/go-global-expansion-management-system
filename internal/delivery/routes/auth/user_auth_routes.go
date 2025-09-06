package routes

import (
	"github.com/gin-gonic/gin"
	AuthHandler "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
)

func RegisterUserAuthRoutes(rg *gin.RouterGroup, authHandler *AuthHandler.UserAuthHandler) {
	auth := rg.Group("/user/auth")

	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
        auth.POST("/logout", middlewares.JWTAuth(), authHandler.Logout)
	}
}
