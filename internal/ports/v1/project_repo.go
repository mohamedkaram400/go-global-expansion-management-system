package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type ProjectRepository interface {
	InsertProject(ctx context.Context, project *entities.Project) (*entities.Project, error)
    FindProjectByID(ctx context.Context, projectID uint) (*entities.Project, error)
    GetAllProjects(ctx context.Context, skip int, limit int) ([]entities.Project, error)
    UpdateProjectByID(ctx context.Context, projectID uint, updates map[string]interface{}) (*entities.Project, error)
    DeleteProjectByID(ctx context.Context, projectID uint) (int, error)
}