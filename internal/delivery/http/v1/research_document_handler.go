package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type ResearchDocumentHandler struct {
	ResearchDocumentService *services.ResearchDocumentService
}

func NewResearchDocumentHandler(researchDocumentService *services.ResearchDocumentService) *ResearchDocumentHandler {
	return &ResearchDocumentHandler{ResearchDocumentService: researchDocumentService}
}

func (h *ResearchDocumentHandler) UploadDocument(c *gin.Context) {

	// Validation for document fields 
	var req requests.UploadDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	savedDoc, err := h.ResearchDocumentService.UploadDocument(c, doc)
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
	searchTerm, err := c.Query("search")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid search term"})
		return
	}

	// Pass search param to service
	documents := h.ResearchDocumentService.SearchOnDocument(c, searchTerm)

	response := generic_api_response.APIResponse{
		Message: "Documents returned successfully",
		Data:    responses.FormatDocuments(documents),
	}

	c.JSON(http.StatusOK, response)
}

