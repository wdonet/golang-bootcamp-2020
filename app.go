package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Todo struct (Model)
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	IsDeleted bool   `json:"isDeleted"`
}

var todos []Todo
var filename string = "./data/todos.csv"

func openForReadCsvFile(filename string) *os.File {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Println("Unable to open for read CSV file!", err)
		return nil
	}
	// r := csv.NewReader(csvfile)
	// r.TrimLeadingSpace = true
	return csvfile
}

func openForWriteCsvFile(filename string) *os.File {
	csvfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		log.Println("Unable to open for write CSV file!", err)
		return nil
	}
	return csvfile
}

// GET /todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// GET /todos/{id}
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Todo{})
}

// POST /todos
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	csvfile := openForWriteCsvFile(filename)
	// Add to the CSV file
	if csvfile == nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Unable to open datasource"}`)
		return
	}
	cw := csv.NewWriter(csvfile)
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	id := strconv.Itoa(todo.ID)
	isDeleted := strconv.FormatBool(todo.IsDeleted)
	if err := cw.Write([]string{id, todo.Task, todo.Status, isDeleted}); err != nil {
		log.Fatalln("Error persisting into csv", filename, err)
	} else {
		cw.Flush()
	}
	csvfile.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func softDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].IsDeleted = true
			todo = todos[idx]
			// todos = append(todos[:idx], todos[idx+1:]...) // hard delete
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func markTodoDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Status = "done"
			todo = todos[idx]
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func markTodoPending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Status = "pending"
			todo = todos[idx]
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Task = params["task"]
			todo = todos[idx]
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func main() {
	router := mux.NewRouter()

	// TODO get data from DB or CSV
	// todos = append(todos, Todo{ID: 10, Task: "Wash dishes", Status: "pending", IsDeleted: false})
	// todos = append(todos, Todo{ID: 20, Task: "Make report", Status: "pending", IsDeleted: false})
	// GET data from CSV
	csvfile := openForReadCsvFile(filename)
	log.Println("Loading records from ", filename)
	r := csv.NewReader(csvfile)
	r.TrimLeadingSpace = true
	var numOfRecords int = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			log.Println("CSV reading done.")
			break
		}
		if err != nil {
			switch t := err.(type) {
			default:
				log.Println("When reading CSV", err)
			case *csv.ParseError:
				log.Println("Ignoring record #", numOfRecords, t)
				continue
			}
		}
		numOfRecords++
		log.Println(record[0], record[1], record[2], record[3])
		// skip headers
		if numOfRecords != 1 {
			id, idErr := strconv.Atoi(record[0])
			isDeleted, delErr := strconv.ParseBool(record[3])
			if idErr == nil && delErr == nil {
				todos = append(todos, Todo{ID: id, Task: record[1], Status: record[2], IsDeleted: isDeleted})
			} else {
				log.Println("Ignoring record #", numOfRecords, idErr, delErr)
			}
		}
	}
	csvfile.Close()

	// Routes
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}/done", markTodoDone).Methods("PUT")
	router.HandleFunc("/todos/{id}/pending", markTodoPending).Methods("PUT")
	router.HandleFunc("/todos/{id}/{task}", updateTask).Methods("PUT")
	router.HandleFunc("/todos/{id}", softDeleteTodo).Methods("DELETE")

	// Start server
	log.Println("Starting server at port [3000].")
	log.Fatal(http.ListenAndServe(":3000", router))
}
