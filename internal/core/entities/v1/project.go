package entities

import (
	"gorm.io/datatypes"
	"time"
)

type Project struct {
	ID             int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Country        string         `json:"country" gorm:"column:country;not null"`
	ServicesNeeded datatypes.JSON `json:"services_needed" gorm:"column:services_needed;not null"`
	Budget         float64        `json:"budget" gorm:"column:budget;not null"`
	Status         string         `gorm:"type:varchar(20);default:'active'"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`

	ClientID int    `json:"client_id" gorm:"not null"`
	Client   Client `gorm:"foreignKey:ClientID;references:ID"`
}
