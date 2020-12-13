package str

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
