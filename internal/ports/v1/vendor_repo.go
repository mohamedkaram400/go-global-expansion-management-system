package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type VendorRepository interface {
	InsertVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error)
    FindVendorByID(ctx context.Context, vendorID int) (*entities.Vendor, error)
    GetAllVendors(ctx context.Context, skip int, limit int) ([]entities.Vendor, error)
    UpdateVendorByID(ctx context.Context, vendorID int, updates map[string]interface{}) (*entities.Vendor, error)
    DeleteVendorByID(ctx context.Context, vendorID int) (int, error)
    FlagExpiredSLAs(ctx context.Context) error
}