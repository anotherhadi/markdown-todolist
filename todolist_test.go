package todolist

import (
	"testing"
)

func TestRead(t *testing.T) {
	todo := New("./testfiles/todolist.md")
	err := todo.Read()
	if err != nil {
		t.Error(err)
	}

	printTodolist(todo)

	if todo.Title != "Reminders" {
		t.Error("Title is not correct")
	}
	if len(todo.Lists) != 2 {
		t.Error("Lists length is not correct")
	}
	if todo.Lists[0].Title != "Shopping" {
		t.Error("List 1 title is not correct")
	}
	if todo.Lists[0].Todos[2].States != Completed {
		t.Error("List 1 todo 3 state is not correct")
	}
}
