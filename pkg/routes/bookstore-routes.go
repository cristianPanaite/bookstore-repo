package routes

import (
	"github.com/cristianPanaite/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("Post")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("Put")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("Delete")
}
