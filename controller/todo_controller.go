package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wdonet/golang-bootcamp-2020/data/handler"
	"github.com/wdonet/golang-bootcamp-2020/domain/model"
	"github.com/wdonet/golang-bootcamp-2020/utilities"
)

// GetTodos process GET on /todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	todos := handler.GetTodosFromFile()
	json.NewEncoder(w).Encode(todos)
}

// GetTodo process GET on /todos/{id}
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	todos := handler.GetTodosFromFile()
	found := utilities.FindTodoByID(params["id"], todos)
	if found == nil {
		found = &model.Todo{}
	}
	json.NewEncoder(w).Encode(found)
}

// CreateTodo process POST on /todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	var todo *model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	if err := handler.SaveTodo(todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Error persisting into csv"}`)
		// log.Fatalln("Error persisting into csv", filename, err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	}
}

// MarkTodoDone process PUT on /todos/{id}/done
func MarkTodoDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	todos := handler.GetTodosFromFile()
	found := utilities.FindTodoByID(params["id"], todos)
	found.Status = utilities.DONE

	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}

// MarkTodoPending process PUT on /todos/{id}/pending
func MarkTodoPending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	todos := handler.GetTodosFromFile()
	found := utilities.FindTodoByID(params["id"], todos)
	found.Status = utilities.PENDING

	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil { // HERE
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}

// UpdateTask process PUT on /todos/{id}/{task}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	todos := handler.GetTodosFromFile()
	found := utilities.FindTodoByID(params["id"], todos)
	found.Task = params["task"]
	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}

// SoftDeleteTodo process DELETE on /todos/{id}
func SoftDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	todos := handler.GetTodosFromFile()
	found := utilities.FindTodoByID(params["id"], todos)
	found.IsDeleted = true
	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
