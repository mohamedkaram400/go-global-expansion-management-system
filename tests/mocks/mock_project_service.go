package mocks

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
)


type MockProjectService struct {
	GetAllProjectsFunc       func(ctx context.Context, skip, limit int) ([]*entities.Project, error)
	GetByEmailFunc           func(ctx context.Context, email string) (*entities.Project, error)
	InsertProjectFunc        func(ctx context.Context, request *requests.ProjectRequest) (*entities.Project, error)
	FindProjectByIDFunc      func(ctx context.Context, id int) (*entities.Project, error)
	UpdateProjectByIDFunc    func(ctx context.Context, id int, updates *entities.Project) (*entities.Project, error)
	DeleteProjectByIDFunc    func(ctx context.Context, id int) (int, error)
	GetProjectCountriesFunc  func(ctx context.Context) (map[int]string, error)
	GetAllActiveProjectsFunc func(ctx context.Context) ([]*entities.Project, error)
}

func (m *MockProjectService) InsertProject(ctx context.Context, request *requests.ProjectRequest) (*entities.Project, error) {
	return m.InsertProjectFunc(ctx, request)
}

func (m *MockProjectService) GetAllProjects(ctx context.Context, skip, limit int) ([]*entities.Project, error) {
	return m.GetAllProjectsFunc(ctx, skip, limit)
}

func (m *MockProjectService) FindProjectByID(ctx context.Context, id int) (*entities.Project, error) {
	return m.FindProjectByIDFunc(ctx, id)
}

func (m *MockProjectService) UpdateProjectByID(ctx context.Context, id int, updates *entities.Project) (*entities.Project, error) {
	return m.UpdateProjectByIDFunc(ctx, id, updates)
}

func (m *MockProjectService) DeleteProjectByID(ctx context.Context, id int) (int, error) {
	return m.DeleteProjectByIDFunc(ctx, id)
}

func (m *MockProjectService) GetProjectCountries(ctx context.Context) (map[int]string, error) {
	return m.GetProjectCountriesFunc(ctx)
}

func (m *MockProjectService) GetAllActiveProjects(ctx context.Context) ([]*entities.Project, error) {
	return m.GetAllActiveProjectsFunc(ctx)
}
