package mocks

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

// MockProjectRepo simulates ProjectRepository behavior
type MockProjectRepo struct {
	GetAllProjectsFunc       func(ctx context.Context, skip, limit int) ([]*entities.Project, error)
	GetByEmailFunc           func(ctx context.Context, email string) (*entities.Project, error)
	InsertProjectFunc        func(ctx context.Context, project *entities.Project) (*entities.Project, error)
	FindProjectByIDFunc      func(ctx context.Context, id int) (*entities.Project, error)
	UpdateProjectByIDFunc    func(ctx context.Context, id int, updates map[string]interface{}) (*entities.Project, error)
	DeleteProjectByIDFunc    func(ctx context.Context, id int) (int, error)
	GetProjectCountriesFunc  func(ctx context.Context) (map[int]string, error)
	GetAllActiveProjectsFunc func(ctx context.Context) ([]*entities.Project, error)
}

func (m *MockProjectRepo) InsertProject(ctx context.Context, Project *entities.Project) (*entities.Project, error) {
	return m.InsertProjectFunc(ctx, Project)
}

func (m *MockProjectRepo) GetAllProjects(ctx context.Context, skip, limit int) ([]*entities.Project, error) {
	return m.GetAllProjectsFunc(ctx, skip, limit)
}

func (m *MockProjectRepo) FindProjectByID(ctx context.Context, id int) (*entities.Project, error) {
	return m.FindProjectByIDFunc(ctx, id)
}

func (m *MockProjectRepo) UpdateProjectByID(ctx context.Context, id int, updates map[string]interface{}) (*entities.Project, error) {
	return m.UpdateProjectByIDFunc(ctx, id, updates)
}

func (m *MockProjectRepo) DeleteProjectByID(ctx context.Context, id int) (int, error) {
	return m.DeleteProjectByIDFunc(ctx, id)
}

func (m *MockProjectRepo) GetProjectCountries(ctx context.Context) (map[int]string, error) {
	return m.GetProjectCountriesFunc(ctx)
}

func (m *MockProjectRepo) GetAllActiveProjects(ctx context.Context) ([]*entities.Project, error) {
	return m.GetAllActiveProjectsFunc(ctx)
}
