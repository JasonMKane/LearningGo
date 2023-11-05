package main

import (
	"fmt"
	"os"
	"todo/database"
)

func main() {
	todo := database.GetItem(1)
	fmt.Println(todo)

	cmdArgs := os.Args[1:]

	fmt.Println(cmdArgs)
	// todo.Name = "This is another new name for item 1"
	// newToDo := database.UpdateToDoItem(todo)
	// println(newToDo)

	// newItem := database.CreateNewToDoItem(*models.NewToDoItem(-1, "Feed Mouse", "This is a new item for Mouse."))
	// fmt.Println(newItem)

	database.GetOpenToDoItems()
}

func listAllOpenItems() {
	fmt.Println(database.GetOpenToDoItems())
}
