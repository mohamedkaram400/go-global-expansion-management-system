package entities

import "time"

type User struct {
	ID           uint      `json:"id" 				gorm:"primaryKey;autoIncrement"`
	Name  		 string    `json:"name" 			gorm:"column:name;not null"`
	Email 		 string    `json:"contact_email" 	gorm:"column:contact_email;unique;not null"`
	Role     	 string    `json:"role" 		    gorm:"column:role;not null"`
	Password     string    `json:"password" 		gorm:"column:password;not null"`
	CreatedAt    time.Time `json:"created_at" 		gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" 		gorm:"column:updated_at;autoUpdateTime"`
} 