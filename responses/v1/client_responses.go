package responses

import "time"


type ClientResponse struct {
    ID           uint   `json:"id"`
    CompanyName  string `json:"company_name"`
    ContactEmail string `json:"contact_email"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}