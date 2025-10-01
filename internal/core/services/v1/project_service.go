package services

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
)

type ProjectService struct {
	Repo ports.ProjectRepository 
}

func NewProjectService(repo ports.ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (svc *ProjectService) GetAllProjects(ctx context.Context, skip int, limit int) ([]entities.Project, error) {
	return svc.Repo.GetAllProjects(ctx, skip, limit)
}
 
func (svc *ProjectService) FindProjectByID(ctx context.Context, projectID uint) (*entities.Project, error) {
	return svc.Repo.FindProjectByID(ctx, projectID)

}

func (svc *ProjectService) InsertProject(ctx context.Context, req *requests.ProjectRequest) (*entities.Project, error) {
	services, _ := json.Marshal(req.ServicesNeeded)
    status := req.Status
    if status == "" { 
        status = entities.ProjectActive
    }

	project := &entities.Project{
		ClientId:				req.ClientId,
		Country:				req.Country,
		Budget:					req.Budget,
		ServicesNeeded: 		services,
		Status: 				status,

	}
	return svc.Repo.InsertProject(ctx, project)
}

func (svc *ProjectService) UpdateProjectByID(ctx context.Context, projectID uint, newProject *entities.Project) (*entities.Project, error) {
	updates := map[string]interface{}{}

    if newProject.Country != "" {
        updates["country"] = newProject.Country
    }

    if newProject.Budget != 0 {
        updates["budget"] = newProject.Budget
    }

    if len(newProject.ServicesNeeded) > 0 {
        updates["services_needed"] = newProject.ServicesNeeded
    }

    if newProject.Status != "" {
        updates["status"] = newProject.Status
    }

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	return svc.Repo.UpdateProjectByID(ctx, projectID, updates)
}

func (svc *ProjectService) DeleteProjectByID(ctx context.Context, projectID uint) (int, error) {
	return svc.Repo.DeleteProjectByID(ctx, projectID)
}

func (svc *ProjectService) FindActiveProjects(ctx context.Context) ([]*entities.Project, error) {
	return svc.Repo.GetAllActiveProjects(ctx)
}