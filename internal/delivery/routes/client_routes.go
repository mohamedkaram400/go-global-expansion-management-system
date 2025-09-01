package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
)

func RegisterClientRoutes(rg *gin.RouterGroup, clientHandler *http.ClientHandler) {
	client := rg.Group("/client")
	client.Use(middlewares.JWTAuth())

	{
		client.POST("/create-client", clientHandler.Create)
		client.GET("/all-clients", clientHandler.Index)
        client.GET("/show-client/:id", clientHandler.Show)
        client.PUT("/update-client/:id", clientHandler.Update)
        client.DELETE("/delete-client/:id", clientHandler.Destroy)
	}
}
