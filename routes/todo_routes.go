package routes

import (
	"github.com/gorilla/mux"
	"github.com/wdonet/golang-bootcamp-2020/controller"
)

// DefineTodoRoutes is the TODO routes definition
func DefineTodoRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/todos").Subrouter()

	// GET
	subrouter.HandleFunc("", controller.GetTodos).Methods("GET")
	subrouter.HandleFunc("/{id}", controller.GetTodo).Methods("GET")

	// POST
	subrouter.HandleFunc("", controller.CreateTodo).Methods("POST")

	// PUT
	subrouter.HandleFunc("/{id}/done", controller.MarkTodoDone).Methods("PUT")
	subrouter.HandleFunc("/{id}/pending", controller.MarkTodoPending).Methods("PUT")
	subrouter.HandleFunc("/{id}/{task}", controller.UpdateTask).Methods("PUT")

	// DELETE
	subrouter.HandleFunc("/{id}", controller.SoftDeleteTodo).Methods("DELETE")
}
