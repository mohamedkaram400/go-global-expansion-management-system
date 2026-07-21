package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/tests/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestFeatchAllVendors(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while fetching vendors", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			GetAllVendorsFunc: func(ctx context.Context, skip int, limit int) ([]*entities.Vendor, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewVendorService(mockRepo)
		vendors, err := svc.GetAllVendors(ctx, 0, 10)

		// Assertions
		assert.Nil(t, vendors)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("successfully fetched", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			GetAllVendorsFunc: func(ctx context.Context, skip int, limit int) ([]*entities.Vendor, error) {
				return seedVendors(t, "Keyo", 2), nil
			},
		}

		svc := services.NewVendorService(mockRepo)
		vendors, err := svc.GetAllVendors(ctx, 0, 10)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, vendors)
		assert.Equal(t, 2, len(vendors))
		assert.Equal(t, "Keyo_1", vendors[0].Name)
	})
}

func TestFindVendorByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: vendor id not found", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			FindVendorByIDFunc: func(ctx context.Context, vendorID int) (*entities.Vendor, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewVendorService(mockRepo)

		vendor, err := svc.FindVendorByID(ctx, 1)

		// Assertions
		assert.Nil(t, vendor)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully returned", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			FindVendorByIDFunc: func(ctx context.Context, vendorID int) (*entities.Vendor, error) {
				return seedVendor(t, "ITWorks"), nil
			},
		}

		svc := services.NewVendorService(mockRepo)

		vendor, err := svc.FindVendorByID(ctx, 1)

		// Assertions
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, vendor)
		assert.Equal(t, "ITWorks", vendor.Name)
	})
}

func TestInsertVendor(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while inserting new vendor", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			InsertVendorFunc: func(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {
				return nil, errors.New("database error")
			},
		}

		servicesArr := []string{"legal", "hiring"}
		countriesArr := []string{"USA", "KSA"}

		svc := services.NewVendorService(mockRepo)
		req := &requests.VendorRequest{Name: "ITworks", CountriesSupported: countriesArr, ServicesOffered: servicesArr, ResponseSlaHours: 3, Rating: 4.5}

		vendor, err := svc.InsertVendor(ctx, req)

		// Assertions
		assert.Nil(t, vendor)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully inserted", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			InsertVendorFunc: func(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {
				vendor.ID = 42
				return vendor, nil
			},
		}

		servicesArr := []string{"legal", "hiring"}
		countriesArr := []string{"USA", "KSA"}

		svc := services.NewVendorService(mockRepo)
		req := &requests.VendorRequest{Name: "ITworks", CountriesSupported: countriesArr, ServicesOffered: servicesArr, ResponseSlaHours: 3, Rating: 4.5}

		vendor, err := svc.InsertVendor(ctx, req)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 42, vendor.ID)
	})
}

func TestUpdateVendorByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while updating new vendor", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			UpdateVendorByIDFunc: func(ctx context.Context, vendorID int, updates map[string]interface{}) (*entities.Vendor, error) {
				return nil, errors.New("database error")
			},
		}

		countriesJSON := convertCountriesSliceToJSON(t)
		servicesJSON := convertServicesSliceToJSON(t)

		svc := services.NewVendorService(mockRepo)
		newVendor := &entities.Vendor{
			Name:               "NewCo",
			ServicesOffered:    datatypes.JSON(servicesJSON),
			CountriesSupported: datatypes.JSON(countriesJSON),
			Rating:             4.7,
		}

		vendor, err := svc.UpdateVendorByID(ctx, 50, newVendor)

		// Assertions
		assert.Nil(t, vendor)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully updated", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			UpdateVendorByIDFunc: func(ctx context.Context, vendorID int, updates map[string]interface{}) (*entities.Vendor, error) {
				return &entities.Vendor{
					ID:     vendorID,
					Name:   updates["name"].(string),
					Rating: updates["rating"].(float64),
				}, nil
			},
		}

		countriesJSON := convertCountriesSliceToJSON(t)
		servicesJSON := convertServicesSliceToJSON(t)

		svc := services.NewVendorService(mockRepo)
		newVendor := &entities.Vendor{
			Name:               "NewCo",
			ServicesOffered:    datatypes.JSON(servicesJSON),
			CountriesSupported: datatypes.JSON(countriesJSON),
			Rating:             4.7,
		}

		vendor, err := svc.UpdateVendorByID(ctx, 50, newVendor)

		assert.NoError(t, err)
		assert.NotNil(t, vendor)
		assert.Equal(t, 50, vendor.ID)
		assert.Equal(t, "NewCo", vendor.Name)
		assert.Equal(t, 4.7, vendor.Rating)
	})
}

func TestDeleteVendorByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: vendor id not found", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			DeleteVendorByIDFunc: func(ctx context.Context, vendorID int) (int, error) {
				return 0, errors.New("database error")
			},
		}

		svc := services.NewVendorService(mockRepo)
		isDeleted, err := svc.DeleteVendorByID(ctx, 1)

		// Assertions
		assert.Equal(t, isDeleted, 0)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully deleted", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			DeleteVendorByIDFunc: func(ctx context.Context, vendorID int) (int, error) {
				return 5, nil
			},
		}

		svc := services.NewVendorService(mockRepo)
		isDeleted, err := svc.DeleteVendorByID(ctx, 5)

		// Assertions
		assert.Equal(t, isDeleted, 5)
		assert.NoError(t, err)
	})
}

func TestFlagExpiredSLAs(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while flagging expired SLAs", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			FlagExpiredSLAsFunc: func(ctx context.Context) error {
				return errors.New("database error")
			},
		}

		svc := services.NewVendorService(mockRepo)

		err := svc.FlagExpiredSLAs(ctx)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully flagged expired SLAs", func(t *testing.T) {
		mockRepo := &mocks.MockVendorRepo{
			FlagExpiredSLAsFunc: func(ctx context.Context) error {
				return nil
			},
		}

		svc := services.NewVendorService(mockRepo)
		err := svc.FlagExpiredSLAs(ctx)

		// Assertions
		assert.NoError(t, err)
	})
}

func seedVendors(t *testing.T, name string, count int) []*entities.Vendor {
	countriesJSON := convertCountriesSliceToJSON(t)
	servicesJSON := convertServicesSliceToJSON(t)

	vendors := make([]*entities.Vendor, 0, count)

	for i := 0; i < count; i++ {
		vendor := &entities.Vendor{
			Name:               fmt.Sprintf("%s_%d", name, i+1),
			CountriesSupported: datatypes.JSON(countriesJSON),
			ServicesOffered:    datatypes.JSON(servicesJSON),
			Rating:             4.5,
			ResponseSlaHours:   3,
		}
		vendors = append(vendors, vendor)
	}

	return vendors
}

func seedVendor(t *testing.T, name string) *entities.Vendor {
	countriesJSON := convertCountriesSliceToJSON(t)
	servicesJSON := convertServicesSliceToJSON(t)

	vendor := &entities.Vendor{
		Name:               name,
		CountriesSupported: datatypes.JSON(countriesJSON),
		ServicesOffered:    datatypes.JSON(servicesJSON),
		Rating:             4.5,
		ResponseSlaHours:   3,
	}

	return vendor
}

func convertCountriesSliceToJSON(t *testing.T) []byte {

	countries := []string{"USA", "KSA"}

	countriesJSON, err := json.Marshal(countries)
	if err != nil {
		t.Fatalf("failed to marshal countries: %v", err)
	}

	return countriesJSON
}

func convertServicesSliceToJSON(t *testing.T) []byte {

	services := []string{"legal", "hiring"}

	servicesJSON, err := json.Marshal(services)
	if err != nil {
		t.Fatalf("failed to marshal services: %v", err)
	}

	return servicesJSON
}
