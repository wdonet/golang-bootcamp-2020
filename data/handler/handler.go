package handler

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/wdonet/golang-bootcamp-2020/domain/model"
	"github.com/wdonet/golang-bootcamp-2020/utilities"
)

const filename string = "./data/todos.csv"

// GetTodosFromFile get data from files as Todos
func GetTodosFromFile() []*model.Todo {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Println("Unable to open for read CSV file!", err)
		return nil
	}
	r := csv.NewReader(csvfile)
	r.TrimLeadingSpace = true

	var numOfRecords int = 0
	var todos []*model.Todo
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
				todos = append(todos, &model.Todo{ID: id, Task: record[1], Status: record[2], IsDeleted: isDeleted})
			} else {
				log.Println("Ignoring record #", numOfRecords, idErr, delErr)
			}
		}
	}
	csvfile.Close()
	return todos
}

// SaveTodo saves a new todo at end of file
func SaveTodo(todo *model.Todo) error {
	csvfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		log.Println("Unable to open for write CSV file!", err)
		return err
	}
	cw := csv.NewWriter(csvfile)
	if err := cw.Write(utilities.ToArrayOfValues(todo)); err != nil {
		log.Fatalln("Error persisting into csv", filename, err)
		csvfile.Close()
		return err
	}
	cw.Flush()
	csvfile.Close()
	return nil
}

func cleanupCsvFile() (*csv.Writer, *os.File, error) {
	csvfile, err := os.Create(filename)
	if err != nil {
		log.Println("Unable to create", filename, err)
		return nil, nil, err
	}
	cw := csv.NewWriter(csvfile)
	if err := cw.Write([]string{"ID", "Task", "Status", "IsDeleted"}); err != nil {
		log.Fatalln("Unable to initialize", filename, err)
		return nil, nil, err
	}
	return cw, csvfile, nil
}

// WriteOnFileAsNew is overwritten any value at the file to replace them with the provided list
func WriteOnFileAsNew(todos []*model.Todo) error {
	csvw, csvfile, err := cleanupCsvFile()
	if err != nil {
		return err
	}
	for _, item := range todos {
		if err := csvw.Write(utilities.ToArrayOfValues(item)); err != nil {
			log.Fatalln("Unable to persist:", item, err)
		}
	}
	csvw.Flush()
	csvfile.Close()
	return nil
}
