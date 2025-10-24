package main

import (
	bookController "librarymvc/internal/books/controllers"
	loanController "librarymvc/internal/loans/controllers"
	userController "librarymvc/internal/users/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	booksController := bookController.NewBooksController()
	usersController := userController.NewUserController()
	loansController := loanController.NewLoanController()

	booksController.RegisterRoutes(router)
	usersController.RegisterRoutes(router)
	loansController.RegisterRouter(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
