package mocks


import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

// MockVendorRepo simulates VendorRepository behavior
type MockVendorRepo struct {
	GetAllVendorsFunc func(ctx context.Context, skip, limit int) ([]*entities.Vendor, error)
	InsertVendorFunc  func(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error)
	FindVendorByIDFunc func(ctx context.Context, id int) (*entities.Vendor, error)
	UpdateVendorByIDFunc func(ctx context.Context, id int, updates map[string]interface{}) (*entities.Vendor, error)
	DeleteVendorByIDFunc func(ctx context.Context, id int) (int, error)
    FlagExpiredSLAsFunc func(ctx context.Context) error

}

func (m *MockVendorRepo) InsertVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {
	return m.InsertVendorFunc(ctx, vendor)
}

func (m *MockVendorRepo) GetAllVendors(ctx context.Context, skip, limit int) ([]*entities.Vendor, error) { 
	return m.GetAllVendorsFunc(ctx, skip, limit)	
}

func (m *MockVendorRepo) FindVendorByID(ctx context.Context, id int) (*entities.Vendor, error) {
	return m.FindVendorByIDFunc(ctx, id)	
}

func (m *MockVendorRepo) UpdateVendorByID(ctx context.Context, id int, updates map[string]interface{}) (*entities.Vendor, error) {
	return m.UpdateVendorByIDFunc(ctx, id, updates)	
}

func (m *MockVendorRepo) DeleteVendorByID(ctx context.Context, id int) (int, error) {
	return m.DeleteVendorByIDFunc(ctx, id)	
}

func (m *MockVendorRepo) FlagExpiredSLAs(ctx context.Context) error {
	return m.FlagExpiredSLAsFunc(ctx)	
}


