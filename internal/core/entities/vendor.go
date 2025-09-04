package entities

import (
	"gorm.io/datatypes"

	"time"
)

type Vendor struct {
	ID                 uint      		`json:"id" 					gorm:"primaryKey;autoIncrement"`
	Name               string    		`json:"name" 				gorm:"column:name;not null"`
	CountriesSupported datatypes.JSON   `json:"countries_supported" gorm:"type:json"`
	ServicesOffered    datatypes.JSON   `json:"services_offered" 	gorm:"type:json"`
	Rating             float64   		`json:"rating" 				gorm:"column:rating;default:0"`
	ResponseSlaHours   uint      		`json:"response_sla_hours" 	gorm:"column:response_sla_hours;default:0"`
	CreatedAt          time.Time 		`json:"created_at" 			gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time 		`json:"updated_at" 			gorm:"column:updated_at;autoUpdateTime"`
}