package auth

import (
	"github.com/gin-gonic/gin"
	AuthHandler "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/auth"
	middlewares "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares/auth"
)

func RegisterUserAuthRoutes(rg *gin.RouterGroup, authHandler *AuthHandler.UserAuthHandler) {
	auth := rg.Group("/user/auth")

	{
		auth.POST("/login", authHandler.Login)
        auth.POST("/logout", middlewares.UserJWTAuth(), authHandler.Logout)
	}
}
