package auth

import (
	"github.com/gin-gonic/gin"
	AuthHandler "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/auth"
	middlewares "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares/auth"
)

func RegisterClientAuthRoutes(rg *gin.RouterGroup, authHandler *AuthHandler.ClientAuthHandler) {
	auth := rg.Group("/client/auth")

	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
        auth.POST("/logout", middlewares.ClientJWTAuth(), authHandler.Logout)
	}
}
