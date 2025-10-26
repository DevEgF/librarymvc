package models

import "time"

type Loan struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	BookId     int64     `gorm:"not null" json:"book_id"`
	UserId     int64     `gorm:"not null" json:"user_id"`
	BorrowAt   time.Time `gorm:"not null" json:"borrow_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `gorm:"not null" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
