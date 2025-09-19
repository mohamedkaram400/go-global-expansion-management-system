package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type MatchHandler struct {
	MatchService *services.MatchService
	ProjectService *services.ProjectService
}

func NewMatchHandler(matchService *services.MatchService) *MatchHandler {
	return &MatchHandler{MatchService: matchService}
}

func (h *MatchHandler) Rebuild(c *gin.Context) {

	idStr := c.Param("project_id")
	// Parse string to uint64 first
    projectID64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
        return
    }

	projectID := uint(projectID64)

	Matchs, err := h.MatchService.Rebuild(c.Request.Context(), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, generic_api_response.APIResponse{
		Message: "Project Matchd successfully",
		Data:    Matchs,
	})
}
