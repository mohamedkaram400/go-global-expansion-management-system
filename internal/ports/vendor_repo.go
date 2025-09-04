package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
)

type VendorRepository interface {
	InsertVendor(ctx context.Context, Vendor *entities.Vendor) (*entities.Vendor, error)
    FindVendorByID(ctx context.Context, VendorID string) (*entities.Vendor, error)
    GetAllVendors(ctx context.Context, skip int, limit int) ([]entities.Vendor, error)
    UpdateVendorByID(ctx context.Context, VendorID string, updates map[string]interface{}) (*entities.Vendor, error)
    DeleteVendorByID(ctx context.Context, VendorID string) (int, error)
}