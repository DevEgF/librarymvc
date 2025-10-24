package controllers

import (
	"librarymvc/internal/loans/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanService models.LoanService
}

func NewLoanController(loanService models.LoanService) *LoanController {
	return &LoanController{loanService: loanService}
}

func (l *LoanController) RegisterRouter(r *gin.Engine) {
	loans := r.Group("/loans")

	{
		loans.POST("", l.CreateLoan)
		loans.GET("/:id", l.GetLoan)
		loans.GET("", l.GetAllLoans)
		loans.PUT("/:id/return", l.ReturnBook)
	}

	user := r.Group("/loans/users")

	{
		user.GET("/:userId/loans", l.GetUserLoans)
	}
}

func (l *LoanController) CreateLoan(ctx *gin.Context) {
	var request struct {
		BookId int64 `json:"book_id"`
		UserId int64 `json:"user_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	loan, err := l.loanService.CreateLoan(request.BookId, request.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (l *LoanController) GetLoan(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	loan, err := l.loanService.GetLoan(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (l *LoanController) GetAllLoans(ctx *gin.Context) {
	loans, err := l.loanService.GetAllLoans()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (l *LoanController) GetUserLoans(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	loans, err := l.loanService.GetUserLoans(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (l *LoanController) ReturnBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	err = l.loanService.ReturnBook(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
