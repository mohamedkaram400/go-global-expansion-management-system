package integrations

import (
	"context"
	"time"
	// "encoding/json"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestGetVendorsForProject(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewMatchRepo(db)
	ctx := context.Background()

	vendors := []entities.Vendor{
		{Name: "VendorUSA", CountriesSupported: datatypes.JSON([]byte(`["USA"]`))},
		{Name: "VendorUAE", CountriesSupported: datatypes.JSON([]byte(`["UAE"]`))},
		{Name: "VendorKSA", CountriesSupported: datatypes.JSON([]byte(`["KSA"]`))},
	}

	for _, v := range vendors {
		db.Create(&v)
	}

	// Create a test project (filter criterion)
	project := &entities.Project{
		Country: "KSA",
	}

	result, err := repo.GetVendorsForProject(ctx, project)

	// t.Logf("Result: %+v", result)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, 1, len(result))
	assert.Equal(t, "VendorKSA", result[0].Name)
}

func TestGetTopVendorsByCountry(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewMatchRepo(db)
	ctx := context.Background()

	// --- Step 1: Prepare sample data ---
	vendors := []entities.Vendor{
		{Name: "VendorA"},
		{Name: "VendorB"},
		{Name: "VendorC"},
	}

	for i := range vendors {
		db.Create(&vendors[i])
	}

	client := entities.Client{CompanyName: "Test Client"}
	db.Create(&client)

	projects := []entities.Project{
		{
			Country:        "KSA",
			ServicesNeeded: datatypes.JSON([]byte(`["IT Consulting"]`)),
			Budget:         10000,
			Status:         "active",
			ClientID:       client.ID,
		},
		{
			Country:        "KSA",
			ServicesNeeded: datatypes.JSON([]byte(`["Cloud Hosting"]`)),
			Budget:         15000,
			Status:         "active",
			ClientID:       client.ID,
		},
		{
			Country:        "UAE",
			ServicesNeeded: datatypes.JSON([]byte(`["Data Analysis"]`)),
			Budget:         20000,
			Status:         "active",
			ClientID:       client.ID,
		},
	}

	for i := range projects {
		db.Create(&projects[i])
	}

	now := time.Now()
	matches := []entities.Match{
		{ProjectID: projects[0].ID, VendorID: vendors[0].ID, Score: 90, CreatedAt: now},
		{ProjectID: projects[0].ID, VendorID: vendors[1].ID, Score: 75, CreatedAt: now},
		{ProjectID: projects[0].ID, VendorID: vendors[2].ID, Score: 70, CreatedAt: now},

		{ProjectID: projects[2].ID, VendorID: vendors[0].ID, Score: 60, CreatedAt: now},
		{ProjectID: projects[2].ID, VendorID: vendors[1].ID, Score: 95, CreatedAt: now},
	}

	for i := range matches {
		db.Create(&matches[i])
	}

	// --- Step 2: Call the repo method ---
	result, err := repo.GetTopVendorsByCountry(ctx, 30)

	// --- Step 3: Debug print ---
	t.Logf("Result: %+v", result)

	// --- Step 4: Assertions ---
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Example check for KSA
	ksaVendors := result["KSA"]
	assert.Equal(t, 3, len(ksaVendors))
	assert.Equal(t, "VendorA", ksaVendors[0].Name)

	// Example check for UAE
	uaeVendors := result["UAE"]
	assert.Equal(t, 2, len(uaeVendors))
	assert.Equal(t, "VendorB", uaeVendors[0].Name)
}

func TestUpsertMatch(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewMatchRepo(db)
	ctx := context.Background()

	// --- Step 1: Prepare related data (client → project → vendor) ---
	client := entities.Client{CompanyName: "Test Client"}
	db.Create(&client)

	project := entities.Project{
		Country:        "KSA",
		ServicesNeeded: datatypes.JSON([]byte(`["Consulting"]`)),
		Budget:         10000,
		Status:         "active",
		ClientID:       client.ID,
	}
	db.Create(&project)

	vendor := entities.Vendor{Name: "VendorA"}
	db.Create(&vendor)

	// --- Step 2: Create a new match ---
	initialMatch := &entities.Match{
		ProjectID: project.ID,
		VendorID:  vendor.ID,
		Score:     90,
	}

	err := repo.UpsertMatch(ctx, initialMatch)

	// t.Logf("Inserted Match: ProjectID=%d, VendorID=%d, Score=%.2f", initialMatch.ProjectID, initialMatch.VendorID, initialMatch.Score)
	// t.Logf("Vendor: ID=%d, Name=%s", vendor.ID, vendor.Name)
	// t.Logf("Project: ID=%d, Country=%s, Budget=%.2f", project.ID, project.Country, project.Budget)

	assert.NoError(t, err, "should insert match without error")

	// --- Step 3: Verify match was inserted ---
	var matchInDB entities.Match
	err = db.First(&matchInDB, "project_id = ? AND vendor_id = ?", project.ID, vendor.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, 90.0, matchInDB.Score)

	// --- Step 4: Call again with same (project_id, vendor_id) but new score ---
	updatedMatch := &entities.Match{
		ProjectID: project.ID,
		VendorID:  vendor.ID,
		Score:     95,
	}

	err = repo.UpsertMatch(ctx, updatedMatch)
	assert.NoError(t, err, "should update existing match")

	// --- Step 5: Verify the score got updated ---
	var updated entities.Match
	err = db.First(&updated, "project_id = ? AND vendor_id = ?", project.ID, vendor.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, 95.0, updated.Score, "score should be updated from 90 to 95")

	// --- Step 6: Sanity check: only one record should exist ---
	var count int64
	db.Model(&entities.Match{}).Count(&count)
	assert.Equal(t, int64(1), count, "should not duplicate match records")
}
