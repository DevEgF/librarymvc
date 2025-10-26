package main

import (
	"librarymvc/internal/database"
	bookModels "librarymvc/internal/books/models"
	loanModels "librarymvc/internal/loans/models"
	userModels "librarymvc/internal/users/models"
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
	if err := database.Connect(); err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	if err := database.DB.AutoMigrate(&bookModels.Book{}, &loanModels.Loan{}, &userModels.User{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	router := gin.Default()

	usersRepository := userRepository.NewUserRepository(database.DB)
	booksRepository := bookRepository.NewBookRepository(database.DB)
	loansRepository := loanRepository.NewLoanRepository(database.DB)

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
