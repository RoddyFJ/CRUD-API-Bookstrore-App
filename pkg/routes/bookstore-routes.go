package routes

import (
	"go-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(routers *mux.Router) {

	routers.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routers.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	routers.HandleFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	routers.HandleFunc("/books/{bookID}", controllers.UpdateBook).Methods("Put")
	routers.HandleFunc("/book/{book}", controllers.DeleteBook).Methods("DELETE")
}
