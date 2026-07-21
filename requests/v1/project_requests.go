package requests

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ProjectRequest struct {
	ClientId       int      `json:"client_id"        validate:"required,gt=0"`
	ServicesNeeded []string `json:"service_needed"   validate:"required,min=1"`
	Country        string   `json:"country"          validate:"required"`
	Budget         float64  `json:"budget"           validate:"required,gte=1"`
	Status         string   `json:"status"           validate:"omitempty,oneof=active completed cancelled"`
}

// Validate performs struct validation
func (r *ProjectRequest) Validate() error {
	return validate.Struct(r)
}

func FormatValidationError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Field() {
			case "ClientId":
				return "client_id is required and must be greater than 0"
			case "ServicesNeeded":
				return "services_needed must contain at least one item"
			case "Country":
				return "country is required"
			case "Budget":
				return "budget must be greater than 0"
			}
		}
	}
	return err.Error()
}
