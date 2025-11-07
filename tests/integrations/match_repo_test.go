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

	t.Logf("Result: %+v", result)
	
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
			Country:         "KSA",
			ServicesNeeded:  datatypes.JSON([]byte(`["IT Consulting"]`)),
			Budget:          10000,
			Status:          "active",
			ClientID:        client.ID,
		},
		{
			Country:         "KSA",
			ServicesNeeded:  datatypes.JSON([]byte(`["Cloud Hosting"]`)),
			Budget:          15000,
			Status:          "active",
			ClientID:        client.ID,
		},
		{
			Country:         "UAE",
			ServicesNeeded:  datatypes.JSON([]byte(`["Data Analysis"]`)),
			Budget:          20000,
			Status:          "active",
			ClientID:        client.ID,
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

