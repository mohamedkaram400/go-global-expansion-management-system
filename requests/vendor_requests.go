package requests

type VendorRequest struct {
	ID                 uint      `json:"id"`
	Name               string    `json:"name"`
	CountriesSupported []string  `json:"countries_supported"`
	ServicesOffered    []string  `json:"services_offered"`
	Rating             float64   `json:"rating"`
	ResponseSlaHours   uint      `json:"response_sla_hours"`
}
