package models

import "time"

type Message struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	From      string    `gorm:"size:50;not null" json:"from"`
	Message   string    `gorm:"size:500;not null" json:"message"`
	To        string    `gorm:"size:50;not null" json:"to"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
