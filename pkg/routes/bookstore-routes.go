package routes

import (
	"github.com/gorilla/mux"
	"github.com/PhucNguyen0304/go-bookstore/pkg/controllers"
	"github.com/PhucNguyen0304/go-bookstore/pkg/middleware"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)
var jwtSecret []byte
var adminEmail string
var adminPassword string

func init() {
    // Load .env file
    err := godotenv.Load("c:/Users/Nguyen Phuc/OneDrive - Thuyloi University/Máy tính/go-bookstore/.env")
    if err != nil {
        fmt.Println("Error loading .env file:", err)
    }

    // Load JWT secret
    jwtSecret = []byte(os.Getenv("JWT_SECRET"))
    if len(jwtSecret) == 0 {
        fmt.Println("JWT_SECRET is not set in the .env file")
    }
}

var BookStoreRoutes = func(router *mux.Router) {
	//Book route
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	//User Route
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
    router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user/", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}",controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/user/{userId}",controllers.UpdateUser).Methods("PUT")

	 // Admin Login
	 router.HandleFunc("/admin/login", controllers.AdminLogin).Methods("POST")

	 router.HandleFunc("/admin/user", middleware.VerifyToken(controllers.GetAllUser, jwtSecret)).Methods("GET")
	 router.HandleFunc("/admin/user/{userId}", middleware.VerifyToken(controllers.GetUserById, jwtSecret)).Methods("GET")

	 // Cart Routes
	 router.HandleFunc("/cart/{userId}/add", controllers.AddItemToCart).Methods("POST")
	 router.HandleFunc("/cart/{userId}", controllers.GetCart).Methods("GET")
}
