package todolist

import (
	"strings"

	"github.com/anotherhadi/markdown"
)

func getTodoState(text string) TodoState {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "- [")
	text = string(text[0])

	if text == string(Completed) {
		return Completed
	} else if text == string(InProgress) {
		return InProgress
	} else if text == string(OnHold) {
		return OnHold
	} else if text == string(Cancelled) {
		return Cancelled
	}

	return NotStarted
}

func (t *TodoList) Read() (err error) {
	err = t.MarkdownFile.Read()
	if err != nil {
		return err
	}

	t.Title = t.MarkdownFile.Title

	for _, section := range t.MarkdownFile.Sections {
		if section.SectionType != markdown.H2 {
			continue
		}
		newlist := List{
			Title: section.Text,
			Todos: []Todo{},
		}
		currentTodo := Todo{}
		for _, line := range section.Lines {
			if line.LineType == markdown.Task {
				if currentTodo.Title != "" {
					newlist.Todos = append(newlist.Todos, currentTodo)
				}
				currentTodo = Todo{
					Title:       line.Text,
					States:      getTodoState(line.Text),
					Description: []string{},
					Settings:    map[string]string{},
				}
				continue
			}
			if currentTodo.Title == "" {
				continue
			}
			if line.LineType == markdown.List {
				if strings.Contains(line.Text, ": ") {
					currentTodo.Settings[strings.Split(line.Text, ": ")[0]] = strings.Split(line.Text, ": ")[1]
				}
			} else {
				currentTodo.Description = append(currentTodo.Description, line.Text)
			}
		}
		if currentTodo.Title != "" {
			newlist.Todos = append(newlist.Todos, currentTodo)
		}
		t.Lists = append(t.Lists, newlist)
	}

	return nil
}
