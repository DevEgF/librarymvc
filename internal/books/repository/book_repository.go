package repository

import (
	"librarymvc/internal/books/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) models.BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) CreateBook(book *models.Book) error {
	return b.db.Create(book).Error
}

func (b *BookRepository) GetBook(id int64) (*models.Book, error) {
	var book models.Book
	if err := b.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *BookRepository) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookRepository) UpdateBook(id int64, book *models.Book) error {
	return b.db.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
}

func (b *BookRepository) DeleteBook(id int64) error {
	return b.db.Delete(&models.Book{}, id).Error
}
