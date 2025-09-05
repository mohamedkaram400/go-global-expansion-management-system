package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
)

func RegisterUserRoutes(rg *gin.RouterGroup, userHandler *http.UserHandler) {
	user := rg.Group("/user")
	user.Use(middlewares.JWTAuth())
	// user.Use(middlewares.AdminAuth())

	{
		user.POST("/create-user", userHandler.Create)
		user.GET("/all-users", userHandler.Index)
        user.GET("/show-user/:id", userHandler.Show)
        user.PUT("/update-user/:id", userHandler.Update)
        user.DELETE("/delete-user/:id", userHandler.Destroy)
	}
}
