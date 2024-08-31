package todolist

import (
	"fmt"
)

func printTodolist(t TodoList) {
	fmt.Println(t.Title)
	for _, list := range t.Lists {
		fmt.Println(list.Title)
		for _, todo := range list.Todos {
			fmt.Println(todo.Title)
			for _, desc := range todo.Description {
				fmt.Println(desc)
			}
		}
	}
}
