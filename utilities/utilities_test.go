package utilities

import (
	"testing"

	"github.com/wdonet/golang-bootcamp-2020/domain/model"
)

func TestTodoObjToArray(t *testing.T) {
	todo := model.Todo{ID: 1, Task: "my task", Status: "pending", IsDeleted: false}
	got := ToArrayOfValues(&todo)
	if got == nil {
		t.Error("ToArrayOfValues should get values")
	}
	if got[0] != "1" {
		t.Errorf("ToArrayOfValues incorrectly got ID [%s]", got[0])
	}
	if got[1] != "my task" {
		t.Errorf("ToArrayOfValues incorrectly got Task [%s]", got[1])
	}
	if got[2] != "pending" {
		t.Errorf("ToArrayOfValues incorrectly got Status [%s]", got[2])
	}
	if got[3] != "false" {
		t.Errorf("ToArrayOfValues incorrectly got IsDeleted flag [%s]", got[3])
	}
}

func TestNilToArray(t *testing.T) {
	if got := ToArrayOfValues(nil); got != nil {
		t.Error("nil should not be converted into anything else than nil")
	}
}

func TestFindTodoByID(t *testing.T) {
	todos := []*model.Todo{
		&model.Todo{ID: 1, Task: "my 1st task", Status: "pending", IsDeleted: false},
		&model.Todo{ID: 2, Task: "my 2nd task", Status: "done", IsDeleted: true},
	}
	got := FindTodoByID("2", todos)
	if got == nil {
		t.Error("FindTodoByID should not return nil")
	}
	if got.ID != 2 {
		t.Error("FindTodoByID found the wrong Todo")
	}
}
