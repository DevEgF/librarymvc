package main

import (
	bookRepository "librarymvc/internal/books/repository"
	loanRepository "librarymvc/internal/loans/repository"
	userRepository "librarymvc/internal/users/repository"

	bookService "librarymvc/internal/books/services"
	loanService "librarymvc/internal/loans/services"
	userService "librarymvc/internal/users/services"

	bookControllers "librarymvc/internal/books/controllers"
	loanControllers "librarymvc/internal/loans/controllers"
	userControllers "librarymvc/internal/users/controllers"
	webControllers "librarymvc/internal/web/controller"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize repositories
	usersRepository := userRepository.NewUserRepository()
	booksRepository := bookRepository.NewBookRepository()
	loansRepository := loanRepository.NewLoanRepository()

	// Initialize services
	usersService := userService.NewUserService(usersRepository)
	booksService := bookService.NewBookService(booksRepository)
	loansService := loanService.NewLoanService(loansRepository, booksService, usersService)

	// Initialize controllers
	booksController := bookControllers.NewBooksController(booksService)
	usersController := userControllers.NewUserController(usersService)
	loansController := loanControllers.NewLoanController(loansService)

	webController := webControllers.NewWebController(booksService, *usersService, *loansService)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// Initialize routers
	booksController.RegisterRoutes(router)
	usersController.RegisterRoutes(router)
	loansController.RegisterRouter(router)
	webController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
