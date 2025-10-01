package entities

import "time"

type VendorStatus struct {
    ID             uint           `gorm:"primaryKey"`
    VendorID       uint           `gorm:"not null"`
    SlaExpired     bool           `gorm:"default:false"`
    CheckRunAt     time.Time      `gorm:"autoCreateTime"`
    LastResponseAt *time.Time     
}