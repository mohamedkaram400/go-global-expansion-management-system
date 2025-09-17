package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
)

func RegisterResearchDocumentRoutes(rg *gin.RouterGroup, ResearchDocumentHandler *http.ResearchDocumentHandler) {
	client := rg.Group("/document")

	{
		client.POST("/research/:project_id/upload", ResearchDocumentHandler.UploadDocument)
		client.GET("/search", ResearchDocumentHandler.SearchOnDocument)
	}
}
