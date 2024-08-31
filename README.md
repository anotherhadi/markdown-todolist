# Markdown todolist

Markdown todolist is a system for creating and managing todolists in markdown.

## Here's the basic idea

- Every markdown file is a different todo list
- The first H1 is the title of the todo list
- H2s are lists
- Task items are list to-do items:
  - Task can have 5 states:
    - `- [ ]` Not started
    - `- [x]` Completed
    - `- [>]` In progress
    - `- [-]` Cancelled
    - `- [?]` On hold
  - The task text (the text after the `- [ ]`) is the todo title
  - The text after the todo title is the todo description/settings:
    - The list items are the event's settings:
      - `dueDate: YYYY-MM-DD`
      - `dueTime: HH:MM`
      - `location: some location`
      - `whateveryouwant: whatever you want`
    - Anything else is the todo description

**Why?** Because I like markdown and I like todolists.
(Also because markdown files are easy to manage with git, and/or with your favorite text editor)

I did the same with calendars, here's the [project](https://github.com/anotherhadi/markdown-calendar)

## The module

I made a simple go module that allows you to read and write markdown todolists.
It's gonna be used on several projects I'm working on, but you can use it too!
Markdown file are parsed with my [markdown writer/reader](https://github.com/anotherhadi/markdown)
