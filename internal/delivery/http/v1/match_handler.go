package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type MatchHandler struct {
	MatchService *services.MatchService
	ProjectService *services.ProjectService
}

func NewMatchHandler(matchService *services.MatchService, projectService *services.ProjectService) *MatchHandler {
	return &MatchHandler{MatchService: matchService, ProjectService: projectService}
}

func (h *MatchHandler) Rebuild(c *gin.Context) {

	projectID, _ := strconv.Atoi(c.Param("project_id"))

	// Load project from DB
	project := entities.Project{}

	project, err := h.ProjectService.FindProjectByID(c, projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	Matchs, err := h.MatchService.Rebuild(c.Request.Context(), project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, generic_api_response.APIResponse{
		Message: "Project Matchd successfully",
		Data:    Matchs,
	})
}
