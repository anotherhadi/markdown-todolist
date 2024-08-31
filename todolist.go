package todolist

import "github.com/anotherhadi/markdown"

type TodoState string

var (
	NotStarted TodoState = " "
	Completed  TodoState = "x"
	InProgress TodoState = ">"
	Cancelled  TodoState = "-"
	OnHold     TodoState = "?"
)

type Todo struct {
	States      TodoState
	Title       string
	Description []string
	Settings    map[string]string
}

type List struct {
	Title string
	Todos []Todo
}

type TodoList struct {
	MarkdownFile markdown.MarkdownFile
	Title        string
	Lists        []List
}

func New(filename string) TodoList {
	md := markdown.New(filename)
	return TodoList{
		MarkdownFile: md,
	}
}
