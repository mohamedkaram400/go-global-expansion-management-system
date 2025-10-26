package integrations

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestGetAllVendors(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewVendorRepo(db)
	ctx := context.Background()

	countriesJSON := convertCountriesSliceToJSON(t)
	servicesJSON := convertServicesSliceToJSON(t)

	// Seed test data
	vendorsData := []entities.Vendor{
		{Name: "Google", CountriesSupported: datatypes.JSON(countriesJSON), ServicesOffered: datatypes.JSON(servicesJSON), Rating: 5.9, ResponseSlaHours: 4},
		{Name: "Microsoft", CountriesSupported: datatypes.JSON(countriesJSON), ServicesOffered: datatypes.JSON(servicesJSON), Rating: 3.7, ResponseSlaHours: 3},
		{Name: "ITWorks", CountriesSupported: datatypes.JSON(countriesJSON), ServicesOffered: datatypes.JSON(servicesJSON), Rating: 4.0, ResponseSlaHours: 5},
	}
	if err := db.Create(&vendorsData).Error; err != nil {
		t.Fatalf("failed to seed vendors: %v", err)
	}

	vendors, err := repo.GetAllVendors(ctx, 0, 3)
	assert.NoError(t, err)
	assert.NotNil(t, vendors)
	assert.Equal(t, 3, len(vendors))
}

func TestGetVendorByID(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewVendorRepo(db)
	ctx := context.Background()

	// Seed one vendor
	vendor := seedVendor(t, db, "ITWorks")

	vendor, err := repo.FindVendorByID(ctx, vendor.ID)
	assert.NoError(t, err)
	assert.NotNil(t, vendor)
	assert.Equal(t, vendor.Name, vendor.Name)
}

func TestInsertVendor(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewVendorRepo(db)
	ctx := context.Background()

	vendor := seedVendor(t, db, "Microsoft")

	// verify record exists
	found, err := repo.FindVendorByID(ctx, vendor.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Microsoft", found.Name)
}

func TestUpdateVendor(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewVendorRepo(db)
	ctx := context.Background()

	vendor := seedVendor(t, db, "Google")

	updates := map[string]interface{}{
		"name":               "ITWorks",
		"rating":             4.9,
		"response_sla_hours": 2,
	}

	updated, err := repo.UpdateVendorByID(ctx, vendor.ID, updates)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.NotZero(t, updated.ID)

	// verify record exists
	found, err := repo.FindVendorByID(ctx, updated.ID)
	assert.NoError(t, err)
	assert.Equal(t, "ITWorks", found.Name)
}

func TestDeleteVendor(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewVendorRepo(db)
	ctx := context.Background()

	vendor := seedVendor(t, db, "Google")

	// 2. Delete the inserted vendor
	rowsAffected, err := repo.DeleteVendorByID(ctx, vendor.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)

	// 3. Verify it's actually deleted
	found, err := repo.FindVendorByID(ctx, vendor.ID)
	assert.Error(t, err) // should return error or not found
	assert.Nil(t, found) // should not find the vendor
}


func seedVendor(t *testing.T, db *gorm.DB, name string) *entities.Vendor {
	countriesJSON := convertCountriesSliceToJSON(t)
	servicesJSON := convertServicesSliceToJSON(t)

	vendor := &entities.Vendor{
		Name:               name,
		CountriesSupported: datatypes.JSON(countriesJSON),
		ServicesOffered:    datatypes.JSON(servicesJSON),
		Rating:             4.5,
		ResponseSlaHours:   3,
	}
	assert.NoError(t, db.Create(&vendor).Error)
	return vendor
}


func convertCountriesSliceToJSON(t *testing.T) ([]byte) {

	countries := []string{"USA", "KSA"}

	countriesJSON, err := json.Marshal(countries)
	if err != nil {
		t.Fatalf("failed to marshal countries: %v", err)
	}

	return countriesJSON
}

func convertServicesSliceToJSON(t *testing.T) ([]byte) {

	services := []string{"legal","hiring"}

	servicesJSON, err := json.Marshal(services)
	if err != nil {
		t.Fatalf("failed to marshal services: %v", err)
	}

	return servicesJSON
}
