package models

import "time"

type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
