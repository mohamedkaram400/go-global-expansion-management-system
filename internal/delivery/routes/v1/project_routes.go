package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
)

func RegisterProjectRoutes(rg *gin.RouterGroup, projectHandler *http.ProjectHandler) {
	project := rg.Group("/project")

	{
		project.POST("/create-project", projectHandler.Create)
		project.GET("/all-projects", projectHandler.Index)
		project.GET("/show-project/:id", projectHandler.Show)
		project.PUT("/update-project/:id", projectHandler.Update)
		project.DELETE("/delete-project/:id", projectHandler.Destroy)
	}
}
