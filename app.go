package main

import (
	"log"
	"net/http"

	"github.com/wdonet/golang-bootcamp-2020/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.DefineTodoRoutes(router)

	// Start server
	log.Println("Starting server at port 3000.")
	log.Fatal(http.ListenAndServe(":3000", router))
}
