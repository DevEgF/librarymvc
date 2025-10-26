package services

import (
	"errors"
	bookService "librarymvc/internal/books/models"
	"librarymvc/internal/loans/models"
	userService "librarymvc/internal/users/models"
	"time"
)

type LoanService struct {
	loanRepo    models.LoanRepository
	bookService bookService.BookService
	userService userService.UserService
}

func NewLoanService(
	loanRepo models.LoanRepository,
	bookService bookService.BookService,
	userService userService.UserService,
) models.LoanService {
	return &LoanService{
		loanRepo:    loanRepo,
		bookService: bookService,
		userService: userService,
	}
}

func (l LoanService) CreateLoan(bookId, userId int64) (*models.Loan, error) {
	book, err := l.bookService.GetBook(bookId)
	if err != nil {
		return nil, err
	}
	if book.Quantity <= 0 {
		return nil, errors.New("the book is not available")
	}

	_, err = l.userService.GetUser(userId)
	if err != nil {
		return nil, err
	}

	activeLoans, err := l.loanRepo.GetActiveUserLoans(userId)
	if err != nil {
		return nil, err
	}

	if len(activeLoans) > 0 {
		return nil, errors.New("user has active loans")
	}

	loan := &models.Loan{
		BookId:    bookId,
		UserId:    userId,
		BorrowAt:  time.Now(),
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = l.loanRepo.CreateLoan(loan)
	if err != nil {
		return nil, err
	}

	book.Quantity--

	if err := l.bookService.UpdateBook(book.ID, book); err != nil {
		return nil, err
	}

	return loan, err
}

func (l LoanService) ReturnBook(loanId int64) error {
	loan, err := l.loanRepo.GetLoan(loanId)
	if err != nil {
		return err
	}

	if loan.Status == "returned" {
		return errors.New("the loan has been returned")
	}

	loan.Status = "returned"
	loan.UpdatedAt = time.Now()
	loan.ReturnedAt = time.Now()

	if err := l.loanRepo.UpdateLoan(loan); err != nil {
		return err
	}

	book, err := l.bookService.GetBook(loan.BookId)
	if err != nil {
		return err
	}

	book.Quantity++
	return l.bookService.UpdateBook(loan.ID, book)
}

func (l LoanService) GetLoan(id int64) (*models.Loan, error) {
	return l.loanRepo.GetLoan(id)
}

func (l LoanService) GetUserLoans(userId int64) ([]*models.Loan, error) {
	return l.loanRepo.GetActiveUserLoans(userId)
}

func (l LoanService) GetAllLoans() ([]*models.Loan, error) {
	return l.loanRepo.GetAllLoans()
}
