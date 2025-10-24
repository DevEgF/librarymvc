package controllers

import (
	"librarymvc/internal/books/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	bookService models.BookService
}

func NewBooksController() *BooksController {
	return &BooksController{}
}

func (b *BooksController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")

	{
		books.POST("", b.CreateBook)
		books.GET("/:id", b.GetBook)
		books.GET("", b.GetAllBooks)
		books.PUT("/:id", b.UpdateBook)
		books.DELETE("/:id", b.DeleteBook)
	}
}

func (b *BooksController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := b.bookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (b *BooksController) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := b.bookService.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (b *BooksController) GetAllBooks(ctx *gin.Context) {
	books, err := b.bookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b *BooksController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = b.bookService.UpdateBook(id, &book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (b *BooksController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
