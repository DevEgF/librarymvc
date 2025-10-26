package models

import "time"

type Book struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title" binding:"required"`
	Author    string    `gorm:"not null" json:"author" binding:"required"`
	Quantity  int       `gorm:"not null" json:"quantity" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
