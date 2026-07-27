package requests

type VendorRequest struct {
	Name               string   `json:"name" binding:"required,min=3,max=100"`
	CountriesSupported []string `json:"countries_supported" binding:"required,dive,required"`
	ServicesOffered    []string `json:"services_offered" binding:"required,dive,required"`
	Rating             float64  `json:"rating" binding:"gte=0,lte=5"`
	ResponseSlaHours   int      `json:"response_sla_hours" binding:"required,gte=1,lte=168"`
}
