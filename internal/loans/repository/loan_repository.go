package repository

import (
	"librarymvc/internal/loans/models"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) models.LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (l *LoanRepository) CreateLoan(loan *models.Loan) error {
	return l.db.Create(loan).Error
}

func (l *LoanRepository) UpdateLoan(loan *models.Loan) error {
	return l.db.Save(loan).Error
}

func (l *LoanRepository) ReturnBook(loanId int64) error {
	return l.db.Model(&models.Loan{}).Where("id = ?", loanId).Update("status", "returned").Error
}

func (l *LoanRepository) GetLoan(id int64) (*models.Loan, error) {
	var loan models.Loan
	if err := l.db.First(&loan, id).Error; err != nil {
		return nil, err
	}
	return &loan, nil
}

func (l *LoanRepository) GetActiveUserLoans(userId int64) ([]*models.Loan, error) {
	var loans []*models.Loan
	if err := l.db.Where("user_id = ? AND status = ?", userId, "active").Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

func (l *LoanRepository) GetAllLoans() ([]*models.Loan, error) {
	var loans []*models.Loan
	if err := l.db.Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}
