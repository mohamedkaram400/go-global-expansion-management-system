package integrations

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"

	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestGetAllProjects(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal", "hiring"}

	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	// Seed test data
	ProjectsData := []*entities.Project{
		{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "cancelled", Budget: 3400},
		{Country: "KSA", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "cancelled", Budget: 3500},
		{Country: "UAE", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "cancelled", Budget: 3600},
	}

	if err := db.Create(&ProjectsData).Error; err != nil {
		t.Fatalf("failed to seed projects: %v", err)
	}

	projects, err := repo.GetAllProjects(ctx, 0, 3)
	assert.NoError(t, err)
	assert.NotNil(t, projects)
	assert.Equal(t, 3, len(projects))
}

func TestGetProjectByID(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal"}
	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	// Seed one project
	projectData := entities.Project{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Budget: 3400}
	if err := db.Create(&projectData).Error; err != nil {
		t.Fatalf("failed to seed project: %v", err)
	}

	project, err := repo.FindProjectByID(ctx, projectData.ID)
	assert.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, projectData.Country, project.Country)
}

func TestInsertProject(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal"}
	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	project := &entities.Project{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Budget: 3400}

	inserted, err := repo.InsertProject(ctx, project)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)
	assert.NotZero(t, inserted.ID)

	// verify record exists
	found, err := repo.FindProjectByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Egypt", found.Country)
}

func TestUpdateProject(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal"}
	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	// 1. Insert a project first to make sure it exists
	project := &entities.Project{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Budget: 3400}

	inserted, err := repo.InsertProject(ctx, project)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)

	updatedServies := []string{"hiring"}
	updatedServicesJSON, err := json.Marshal(updatedServies)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	updates := map[string]interface{}{
		"Country":        "KSA",
		"ServicesNeeded": updatedServicesJSON,
		"Budget":         4000,
	}

	updated, err := repo.UpdateProjectByID(ctx, project.ID, updates)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.NotZero(t, updated.ID)

	// verify record exists
	found, err := repo.FindProjectByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "KSA", found.Country)
}

func TestDeleteProject(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal"}
	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	// 1. Insert a project first to make sure it exists
	project := &entities.Project{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Budget: 3400}

	inserted, err := repo.InsertProject(ctx, project)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)

	// 2. Delete the inserted project
	rowsAffected, err := repo.DeleteProjectByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)

	// 3. Verify it's actually deleted
	found, err := repo.FindProjectByID(ctx, inserted.ID)
	assert.Error(t, err) // should return error or not found
	assert.Nil(t, found) // should not find the project
}

func TestGetProjectCountries(t *testing.T) {
	db := conn.SetupTestDB()
	// db = db.Debug()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal", "hiring"}

	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	// Seed test data
	ProjectsData := []*entities.Project{
		{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3400},
		{Country: "KSA", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3500},
		{Country: "UAE", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3600},
	}

	if err := db.Create(&ProjectsData).Error; err != nil {
		t.Fatalf("failed to seed projects: %v", err)
	}

	projectCountries, err := repo.GetAllActiveProjects(ctx)

	// prettyPrint(projectCountries)

	assert.NoError(t, err)
	assert.NotNil(t, projectCountries)
	assert.Contains(t, []string{"Egypt", "KSA"}, "Egypt")
}

func TestGetAllActiveProjects(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewProjectRepo(db)
	ctx := context.Background()

	servicesNeeded := []string{"legal", "hiring"}

	servicesJSON, err := json.Marshal(servicesNeeded)
	if err != nil {
		t.Fatalf("failed to marshal servicesNeeded: %v", err)
	}

	// Seed test data
	ProjectsData := []*entities.Project{
		{Country: "Egypt", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3400},
		{Country: "KSA", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3500},
		{Country: "UAE", ServicesNeeded: datatypes.JSON(servicesJSON), Status: "active", Budget: 3600},
	}

	if err := db.Create(&ProjectsData).Error; err != nil {
		t.Fatalf("failed to seed projects: %v", err)
	}

	activeProjects, err := repo.GetAllActiveProjects(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, activeProjects)
}

func prettyPrint(data any) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
