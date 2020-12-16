package utilities

import (
	"strconv"

	"github.com/wdonet/golang-bootcamp-2020/domain/model"
)

// ToArrayOfValues convert a Todo into an array of values
func ToArrayOfValues(todo *model.Todo) []string {
	if todo == nil {
		return nil
	}
	id := strconv.Itoa(todo.ID)
	isDeleted := strconv.FormatBool(todo.IsDeleted)
	return []string{id, todo.Task, todo.Status, isDeleted}
}

// FindTodoByID find by id a Todo in the list
func FindTodoByID(idStr string, todos []*model.Todo) *model.Todo {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil
	}
	var todo *model.Todo
	for idx, item := range todos {
		if item.ID == id {
			todo = todos[idx]
			break
		}
	}
	return todo
}
