# Welcome to my small CLI to-do app in Go

Since I live and breathe in the CLI, I have decided to make my seemingly easy life more complicated, in terms of taking notes.

So I decided to make this small to-do app in Golang, that lives in the CLI.

<img width="934" height="226" alt="cli_todo" src="https://github.com/user-attachments/assets/16ea5488-0aa6-497a-bb4f-f72cb6ce60a7" />

## How to use it:

Clone the Repository and Run `go run ./` + `"commands"` to use the app.

#### Commands

- `-add "Your task name"` - Adds a task in the To-do.
- `-edit 1(your note index number):"New_title"` - Edits the existing list title. Simply write your item index number in place of `1`

- `-del 1(your note index number)` - Deletes the note at the given `index`
- `-toggle 1(your note index number)` - Marks the selected index as completed.
