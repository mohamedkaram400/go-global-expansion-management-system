package entities

import "time"

type VendorStatus struct {
    ID             int            `gorm:"primaryKey"`
    SlaExpired     bool           `gorm:"default:false"`
    CheckRunAt     time.Time      `gorm:"autoCreateTime"`
    LastResponseAt *time.Time     

    VendorID int    `gorm:"not null"`
    Vendor   Vendor `gorm:"foreignKey:VendorID;references:ID"`
}