package models

import "time"


type Book struct {
	ID     uint   `gorm:"primaryKey" `
	Title string `gorm:"type:varchar(150);not null"`
	Author string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}