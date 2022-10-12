package models

import "time"

type User struct {
	ID        string    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
