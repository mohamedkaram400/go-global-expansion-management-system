package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
)

func RegisterAnalyticsRoutes(rg *gin.RouterGroup, analyticsHandler *http.AnalyticsHandler) {
	client := rg.Group("/analytics")

	{
		client.GET("/top-vendors", analyticsHandler.GenerateAnalytics)
	}
}
