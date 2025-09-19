package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type ResearchDocumentHandler struct {
	ResearchDocumentService *services.ResearchDocumentService
	ProjectService *services.ProjectService
}

func NewResearchDocumentHandler(researchDocumentService *services.ResearchDocumentService,  projectService *services.ProjectService) *ResearchDocumentHandler {
	return &ResearchDocumentHandler{ResearchDocumentService: researchDocumentService,  ProjectService: projectService}

}

func (h *ResearchDocumentHandler) UploadDocument(c *gin.Context) {
    ctx := c.Request.Context()

	// Validation for document fields 
	var req requests.UploadDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate project exists in MySQL
	_, err := h.ProjectService.FindProjectByID(ctx, req.ProjectId)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Project id not found"})
		return 
	}

	// Prepare document object
	doc := &entities.Document{
		ProjectId: req.ProjectId,
		Title:     req.Title,
		Content:   req.Content,
		Tags:      req.Tags,
	}

	// Pass values to service 
	savedDoc, err := h.ResearchDocumentService.UploadDocument(ctx, doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, generic_api_response.APIResponse{
		Message: "Document uploaded successfully",
		Data:    responses.FormatDocument(savedDoc),
	})
}

func (h *ResearchDocumentHandler) SearchOnDocument(c *gin.Context) {

	// Extract the search param
	searchTerm := c.Query("search")
	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search term required"})
		return
	}

	// Pass search param to service
	documents, err := h.ResearchDocumentService.SearchOnDocument(c, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := generic_api_response.APIResponse{
		Message: "Documents returned successfully",
		Data:    responses.FormatDocuments(documents),
	}

	c.JSON(http.StatusOK, response)
}

