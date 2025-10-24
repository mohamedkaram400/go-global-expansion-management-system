package services

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/tests/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestFeatchAllProjects(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while fetching projects", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			GetAllProjectsFunc: func(ctx context.Context, skip int, limit int) ([]entities.Project, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewProjectService(mockRepo)
		projects, err := svc.GetAllProjects(ctx, 0, 10)

		// Assertions
		assert.Nil(t, projects)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("successfully featched", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			GetAllProjectsFunc: func(ctx context.Context, skip int, limit int) ([]entities.Project, error) {

				servicesNeeded := []string{"legal","hiring"}

				servicesJSON, err := json.Marshal(servicesNeeded)
				if err != nil {
					t.Fatalf("failed to marshal servicesNeeded: %v", err)
				}

				return []entities.Project{
					{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "cancelled", Budget: 3400},
					{Country: "KSA", ServicesNeeded:   datatypes.JSON(servicesJSON), Status: "cancelled", Budget: 3500},
					{Country: "UAE", ServicesNeeded:   datatypes.JSON(servicesJSON), Status: "cancelled", Budget: 3600},
				}, nil
			},
		}

		svc := services.NewProjectService(mockRepo)
		projects, err := svc.GetAllProjects(ctx, 0, 10)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, projects)
		assert.Equal(t, 3, len(projects))
		assert.Equal(t, "Egypt", projects[0].Country)
	})
}

func TestFindProjectByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: project id not found", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			FindProjectByIDFunc: func(ctx context.Context, ProjectID int) (*entities.Project, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewProjectService(mockRepo)

		project, err := svc.FindProjectByID(ctx, 1)

		// Assertions
		assert.Nil(t, project)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully returned", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			FindProjectByIDFunc: func(ctx context.Context, ProjectID int) (*entities.Project, error) {
				return &entities.Project{ID: ProjectID, Country: "Egypt"}, nil
			},
		}

		svc := services.NewProjectService(mockRepo)

		project, err := svc.FindProjectByID(ctx, 1)

		// Assertions
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, "Egypt", project.Country)
	})
}

func TestInsertProject(t *testing.T) {
	ctx := context.Background()

	t.Run("successfully inserted", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			InsertProjectFunc: func(ctx context.Context, project *entities.Project) (*entities.Project, error) {
				project.ID = 42
				return project, nil
			},
		}

		servicesNeeded := []string{"legal","hiring"}
				
		svc := services.NewProjectService(mockRepo)
		req := &requests.ProjectRequest{Country: "Egypt", ServicesNeeded: servicesNeeded, Status: "active", Budget: 3400}

		project, err := svc.InsertProject(ctx, req)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 42, project.ID)
	})
}

func TestUpdateProjectByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: database update failure", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.Project, error) {
				return nil, nil // no conflict
			},
			UpdateProjectByIDFunc: func(ctx context.Context, ProjectID int, updates map[string]interface{}) (*entities.Project, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewProjectService(mockRepo)

		newProject := &entities.Project{Country: "UAE", Budget: 70000.0}	

		project, err := svc.UpdateProjectByID(ctx, 1, newProject)

		assert.Nil(t, project)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully updated", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			UpdateProjectByIDFunc: func(ctx context.Context, ProjectID int, updates map[string]interface{}) (*entities.Project, error) {
				return &entities.Project{
					ID:           42,
					Country:  updates["country"].(string),
					Budget: updates["budget"].(float64),
				}, nil
			},
		}

		svc := services.NewProjectService(mockRepo)
		newProject := &entities.Project{Country: "UAE", Budget: 70000.0}
		project, err := svc.UpdateProjectByID(ctx, 1, newProject)

		assert.NoError(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, 42, project.ID)
		assert.Equal(t, "UAE", project.Country)
		assert.Equal(t, 70000.0, project.Budget)
	})
}

func TestDeleteProjectByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: project id not found", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			DeleteProjectByIDFunc: func(ctx context.Context, ProjectID int) (int, error) {
				return 0, errors.New("database error")
			},
		}

		svc := services.NewProjectService(mockRepo)
		isDeleted, err := svc.DeleteProjectByID(ctx, 1)

		// Assertions
		assert.Equal(t, isDeleted, 0)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully deleted", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			DeleteProjectByIDFunc: func(ctx context.Context, ProjectID int) (int, error) {
				return 1, nil
			},
		}

		svc := services.NewProjectService(mockRepo)
		isDeleted, err := svc.DeleteProjectByID(ctx, 1)

		// Assertions
		assert.Equal(t, isDeleted, 1)
		assert.NoError(t, err)
	})
}

func TestGetAllActiveProjects(t *testing.T) {
	ctx := context.Background()

	t.Run("error: database error", func(t *testing.T) {
		mockRepo := &mocks.MockProjectRepo{
			GetAllActiveProjectsFunc: func(ctx context.Context) ([]*entities.Project, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewProjectService(mockRepo)
		projects, err := svc.GetAllActiveProjects(ctx)

		assert.EqualError(t, err, "database error")
		assert.Nil(t, projects)
	})

	t.Run("active projects returned successfully", func(t *testing.T) {
			mockRepo := &mocks.MockProjectRepo{
			GetAllActiveProjectsFunc: func(ctx context.Context) ([]*entities.Project, error) {

				servicesNeeded := []string{"legal","hiring"}

				servicesJSON, err := json.Marshal(servicesNeeded)
				if err != nil {
					t.Fatalf("failed to marshal servicesNeeded: %v", err)
				}

				return []*entities.Project{
					{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3400},
					{Country: "KSA", ServicesNeeded:   datatypes.JSON(servicesJSON), Status: "active", Budget: 3500},
					{Country: "UAE", ServicesNeeded:   datatypes.JSON(servicesJSON), Status: "active", Budget: 3600},
				}, nil
			},
		}

		svc := services.NewProjectService(mockRepo)
		projects, err := svc.GetAllActiveProjects(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, projects)
	})
}