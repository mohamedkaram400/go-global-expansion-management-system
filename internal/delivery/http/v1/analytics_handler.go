package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type AnalyticsHandler struct {
	AnalyticsService *services.AnalyticsService
}

func NewAnalyticsHandler(AnalyticsService *services.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{AnalyticsService: AnalyticsService}
}

func (h *AnalyticsHandler) GenerateAnalytics(c *gin.Context) {

	analytics, err := h.AnalyticsService.GenerateAnalytics(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	response := generic_api_response.APIResponse{
		Message: "Top vendors analytics retrieved successfully",
		Data: analytics,
	}

	c.JSON(http.StatusOK, response)
}
