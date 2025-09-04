package responses

import (
	"encoding/json"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
)


type VendorResponse struct {
	ID                 uint      `json:"id"`                 
	Name               string    `json:"name"`
	CountriesSupported []string  `json:"countries_supported"`
	ServicesOffered    []string  `json:"services_offered"`
	Rating             float64   `json:"rating"`
	ResponseSlaHours   uint      `json:"response_sla_hours"`
}


func FormatVendor(vendor *entities.Vendor) VendorResponse {
    var countries []string
	var services []string

	// Convert JSON to []string
	_ = json.Unmarshal(vendor.CountriesSupported, &countries)
	_ = json.Unmarshal(vendor.ServicesOffered, &services)

	return VendorResponse{
		ID:                 vendor.ID,
		Name:               vendor.Name,
		CountriesSupported: countries,
		ServicesOffered:    services,
		Rating:             vendor.Rating,
		ResponseSlaHours:   vendor.ResponseSlaHours,
	}
}

func FormatVendors(vendors []entities.Vendor) []VendorResponse {
	responses := make([]VendorResponse, 0, len(vendors))
	for _, v := range vendors {
		responses = append(responses, FormatVendor(&v))
	}
	return responses
}