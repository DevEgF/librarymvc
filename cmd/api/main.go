package main

import (
	bookRepository "librarymvc/internal/books/repository"
	loanRepository "librarymvc/internal/loans/repository"
	userRepository "librarymvc/internal/users/repository"

	bookService "librarymvc/internal/books/services"
	loanService "librarymvc/internal/loans/services"
	userService "librarymvc/internal/users/services"

	bookController "librarymvc/internal/books/controllers"
	loanController "librarymvc/internal/loans/controllers"
	userController "librarymvc/internal/users/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	usersRepository := userRepository.NewUserRepository()
	booksRepository := bookRepository.NewBookRepository()
	loansRepository := loanRepository.NewLoanRepository()

	usersService := userService.NewUserService(usersRepository)
	booksService := bookService.NewBookService(booksRepository)
	loansService := loanService.NewLoanService(loansRepository, booksService, usersService)

	booksController := bookController.NewBooksController(booksService)
	usersController := userController.NewUserController(usersService)
	loansController := loanController.NewLoanController(loansService)

	booksController.RegisterRoutes(router)
	usersController.RegisterRoutes(router)
	loansController.RegisterRouter(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
