package entities

import (
    "time"
	"gorm.io/datatypes"
)

type ProjectStatus string

const (
    ProjectActive    ProjectStatus = "active"
    ProjectCompleted ProjectStatus = "completed"
    ProjectCancelled ProjectStatus = "cancelled"
)

type Project struct {
	ID             uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Country        string          `json:"country" gorm:"column:country;not null"`
	ServicesNeeded datatypes.JSON  `json:"services_needed" gorm:"column:services_needed;not null"`
	Budget         float64         `json:"budget" gorm:"column:budget;not null"`
	Status         ProjectStatus   `json:"status" gorm:"type:enum('active','completed','cancelled');default:'active'"`
	CreatedAt      time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `json:"updated_at" gorm:"autoUpdateTime"`

	ClientID uint   			   `json:"client_id" gorm:"not null"`
    Client   Client 			   `gorm:"foreignKey:ClientID;references:ID"`
}