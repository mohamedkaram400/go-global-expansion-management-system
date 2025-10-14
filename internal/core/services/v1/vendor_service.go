package services

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/datatypes"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
)

type VendorService struct {
	Repo ports.VendorRepository 
}

func NewVendorService(repo ports.VendorRepository) *VendorService {
	return &VendorService{Repo: repo}
}

func (svc *VendorService) GetAllVendors(ctx context.Context, skip int, limit int) ([]entities.Vendor, error) {
	return svc.Repo.GetAllVendors(ctx, skip, limit)
}
 
func (svc *VendorService) FindVendorByID(ctx context.Context, VendorID int) (*entities.Vendor, error) {
	return svc.Repo.FindVendorByID(ctx, VendorID)

}

func (svc *VendorService) InsertVendor(ctx context.Context, req *requests.VendorRequest) (*entities.Vendor, error) {
	countries, _ := json.Marshal(req.CountriesSupported)
	services, _ := json.Marshal(req.ServicesOffered)

	Vendor := &entities.Vendor{
		Name:				req.Name,
		CountriesSupported:	datatypes.JSON(countries),
		ServicesOffered:	datatypes.JSON(services),
		Rating:				req.Rating,
		ResponseSlaHours: 	req.ResponseSlaHours,
	}

	return svc.Repo.InsertVendor(ctx, Vendor)
}

func (svc *VendorService) UpdateVendorByID(ctx context.Context, VendorID int, newVendor *entities.Vendor) (*entities.Vendor, error) {
	updates := map[string]interface{}{}

    if newVendor.Name != "" {
        updates["name"] = newVendor.Name
    }

    if len(newVendor.CountriesSupported) > 0 {
        updates["countries_supported"] = newVendor.CountriesSupported
    }

    if len(newVendor.ServicesOffered) > 0 {
        updates["services_offered"] = newVendor.ServicesOffered
    }

    if newVendor.Rating != 0 {
        updates["rating"] = newVendor.Rating
    }

    if newVendor.ResponseSlaHours != 0 {
        updates["response_sla_hours"] = newVendor.ResponseSlaHours
    }
	

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	return svc.Repo.UpdateVendorByID(ctx, VendorID, updates)
}

func (svc *VendorService) DeleteVendorByID(ctx context.Context, VendorID int) (int, error) {
	return svc.Repo.DeleteVendorByID(ctx, VendorID)

}

func (svc *VendorService) FlagExpiredSLAs(ctx context.Context) error {
	return svc.Repo.FlagExpiredSLAs(ctx)
}