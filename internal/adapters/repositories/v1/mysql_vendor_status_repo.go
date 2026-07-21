package repositories

import (
	"context"
	"time"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"gorm.io/gorm"
)

type VendorStatusRepo struct {
	DB *gorm.DB
}

func NewVendorStatusRepo(db *gorm.DB) *VendorStatusRepo {
	return &VendorStatusRepo{DB: db}
}

func (r *VendorStatusRepo) UpdateLastResponse(ctx context.Context, vendorID int64) error {
	return r.DB.WithContext(ctx).Model(&entities.VendorStatus{}).
		Where("vendor_id = ?", vendorID).
		Update("last_response_at", time.Now()).Error
}
