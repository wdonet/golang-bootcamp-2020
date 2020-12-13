package str

import (
	"strconv"

	"github.com/wdonet/golang-bootcamp-2020/domain/model"
)

// ToArrayOfValues convert a Todo into an array of values
func ToArrayOfValues(todo *model.Todo) []string {
	id := strconv.Itoa(todo.ID)
	isDeleted := strconv.FormatBool(todo.IsDeleted)
	return []string{id, todo.Task, todo.Status, isDeleted}
}
