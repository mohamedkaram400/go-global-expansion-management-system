package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type ProjectHandler struct {
	Service *services.ProjectService
}

func NewProjectHandler(service *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{Service: service} 
}

func (h *ProjectHandler) Index(c *gin.Context) {

    skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    projects, err := h.Service.GetAllProjects(c.Request.Context(), skip, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Projects returned successfully",
		Data:    responses.FormatProjects(projects),
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProjectHandler) Show(c *gin.Context) {
	idStr := c.Param("id")

	// Parse string to int first
    projectID64, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
        return
    }

	projectID := int(projectID64)


	newProject, err := h.Service.FindProjectByID(c.Request.Context(), projectID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Project Returned Successfully",
		Data:    responses.FormatProject(newProject),
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var req requests.ProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": requests.FormatValidationError(err)})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": requests.FormatValidationError(err)})
		return
	}

	newProject, err := h.Service.InsertProject(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := generic_api_response.APIResponse{
		Message: "Project Created Successfully",
		Data:    responses.FormatProject(newProject),
	}

	c.JSON(http.StatusCreated, response)
}

func (h *ProjectHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	// Parse string to int first
    projectID64, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
        return
    }

	projectID := int(projectID64)

    var req requests.ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    services, _ := json.Marshal(req.ServicesNeeded)

	updates := &entities.Project{
		ClientID:				req.ClientId, 
		Country:				req.Country,
		Budget:					req.Budget,
		ServicesNeeded: 		services,
		Status: 				req.Status,
	}

	project, err := h.Service.UpdateProjectByID(c.Request.Context(), projectID, updates)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
        Message: "Project updated successfully",
		Data:    responses.FormatProject(project),
    }

	c.JSON(http.StatusOK, response)
}

func (h *ProjectHandler) Destroy(c *gin.Context) {
	idStr := c.Param("id")

	// Parse string to int first
    projectID64, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
        return
    }

	projectID := int(projectID64)

    _, err = h.Service.DeleteProjectByID(c.Request.Context(), projectID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	

	c.JSON(http.StatusNoContent, nil)
}
