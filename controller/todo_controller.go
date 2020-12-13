package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wdonet/golang-bootcamp-2020/data/handler"
	"github.com/wdonet/golang-bootcamp-2020/domain/model"
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
	params := mux.Vars(r)
	todos := handler.GetTodosFromFile()
	for _, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Todo{})
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
	var todo *model.Todo
	todos := handler.GetTodosFromFile()
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Status = "done"
			todo = todos[idx]
			break
		}
	}
	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// MarkTodoPending process PUT on /todos/{id}/pending
func MarkTodoPending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo *model.Todo
	todos := handler.GetTodosFromFile()
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Status = "pending"
			todo = todos[idx]
			break
		}
	}
	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil { // HERE
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// UpdateTask process PUT on /todos/{id}/{task}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo *model.Todo
	todos := handler.GetTodosFromFile()
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Task = params["task"]
			todo = todos[idx]
			break
		}
	}
	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// SoftDeleteTodo process DELETE on /todos/{id}
func SoftDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo *model.Todo
	todos := handler.GetTodosFromFile()
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].IsDeleted = true
			todo = todos[idx]
			// todos = append(todos[:idx], todos[idx+1:]...) // hard delete
			break
		}
	}
	// Write everything again
	if err := handler.WriteOnFileAsNew(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to update datasource"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}
