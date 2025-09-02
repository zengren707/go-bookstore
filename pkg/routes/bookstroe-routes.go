package routes

import (
	"github.com/gorilla/mux"
	"github.com/zengren707/go-bookstroe/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
