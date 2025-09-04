package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"gorm.io/gorm"
)

type VendorRepo struct {
	DB *gorm.DB
}

func NewVendorRepo(db *gorm.DB) *VendorRepo {
	return &VendorRepo{DB: db}
}

func (r *VendorRepo) GetAllVendors(ctx context.Context, skip int, limit int) ([]entities.Vendor, error) {
	var vendors []entities.Vendor
	if err := r.DB.WithContext(ctx).
		Offset(skip).
		Limit(limit).
		Find(&vendors).Error; err != nil {
		return nil, err
	}
	return vendors, nil
}

func (r *VendorRepo) FindVendorByID(ctx context.Context, vendorID string) (*entities.Vendor, error) {
	var vendor entities.Vendor
	if err := r.DB.WithContext(ctx).
		Where("id = ?", vendorID).
		First(&vendor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	fmt.Printf("✅ Vendor inserted: %+v\n", vendor) // debug

	return &vendor, nil
}

func (r *VendorRepo) InsertVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {
	if err := r.DB.WithContext(ctx).Create(vendor).Error; err != nil {
		return nil, err
	}
	return vendor, nil
}

func (r *VendorRepo) UpdateVendorByID(ctx context.Context, vendorID string, updates map[string]interface{}) (*entities.Vendor, error) {

	vendor := &entities.Vendor{}

	// Update the Vendor
	if err := r.DB.WithContext(ctx).
		Model(vendor).
		Where("id = ?", vendorID).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	// Fetch the updated record
	if err := r.DB.WithContext(ctx).
		Where("id = ?", vendorID).
		First(vendor).Error; err != nil {
		return nil, err
	}

	return vendor, nil
}

func (r *VendorRepo) DeleteVendorByID(ctx context.Context, vendorID string) (int, error) {
	if err := r.DB.WithContext(ctx).
		Where("id = ?", vendorID).
		Delete(&entities.Vendor{}).Error; err != nil {
		return 0, err
	}

	return 1, nil
}
