package models

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name" binding:"required"`
	Email     string    `gorm:"not null;unique" json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
