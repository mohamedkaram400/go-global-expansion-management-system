package entities

import "time"

type Client struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	CompanyName  string    `gorm:"column:company_name;not null"`
	ContactEmail string    `gorm:"column:contact_email;unique;not null"`
	Password     string    `gorm:"column:password;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}