package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
)

type ProjectService interface {
	GetAllProjects(ctx context.Context, skip int, limit int) ([]*entities.Project, error)
	FindProjectByID(ctx context.Context, projectID int) (*entities.Project, error)
    InsertProject(ctx context.Context, req *requests.ProjectRequest) (*entities.Project, error)
    UpdateProjectByID(ctx context.Context, projectID int, newProject *entities.Project) (*entities.Project, error)
    DeleteProjectByID(ctx context.Context, projectID int) (int, error)
    GetAllActiveProjects(ctx context.Context) ([]*entities.Project, error)
}
