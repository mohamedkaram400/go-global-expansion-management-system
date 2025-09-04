package entities

import "time"

type Client struct {
	ID           uint      `json:"id" 				gorm:"primaryKey;autoIncrement"`
	CompanyName  string    `json:"company_name" 	gorm:"column:company_name;not null"`
	ContactEmail string    `json:"contact_email" 	gorm:"column:contact_email;unique;not null"`
	Password     string    `json:"password" 		gorm:"column:password;not null"`
	CreatedAt    time.Time `json:"created_at" 		gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" 		gorm:"column:updated_at;autoUpdateTime"`
} 