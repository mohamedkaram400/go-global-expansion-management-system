package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
)

func RegisterMatchRoutes(rg *gin.RouterGroup, clientHandler *http.MatchHandler) {
	client := rg.Group("/match")

	{
		client.POST("/projects/:id/match/rebuild", clientHandler.Rebuild)
	}
}
