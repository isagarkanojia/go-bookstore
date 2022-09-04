package routes

import (
	"github.com/gorilla/mux"
	"github.com/isagarkanojia/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/books/", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/books/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")

}