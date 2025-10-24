package models

import "time"

type Loan struct {
	ID         int64     `json:"id"`
	BookId     int64     `json:"book_id"`
	UserId     int64     `json:"user_id"`
	BorrowAt   time.Time `json:"borrow_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
