package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"todo/database"
	"todo/models"
)

func main() {
	// todo := database.GetItem(1)
	// fmt.Println(todo)

	cmdArgs := os.Args[1:]

	if slices.Contains(cmdArgs, "list") {
		listAllOpenItems()
	}

	if slices.Contains(cmdArgs, "new") {
		createNewItem()
	}

	if slices.Contains(cmdArgs, "complete") {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("ToDo Item Name: ")
		id, _ := reader.ReadString('\n')
		itemId, _ := strconv.Atoi(id)
		setItemComplete(itemId)
	}

	// todo.Name = "This is another new name for item 1"
	// newToDo := database.UpdateToDoItem(todo)
	// println(newToDo)

	// newItem := database.CreateNewToDoItem(*models.NewToDoItem(-1, "Feed Mouse", "This is a new item for Mouse."))
	// fmt.Println(newItem)

	// database.GetOpenToDoItems()
}

func listAllOpenItems() {
	items := database.GetOpenToDoItems()

	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}

func createNewItem() {
	var name string
	var description string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ToDo Item Name: ")
	name, _ = reader.ReadString('\n')

	fmt.Print("ToDo Item Description: ")
	description, _ = reader.ReadString('\n')

	todoItem := models.ToDoItem{Name: name, Description: description, IsComplete: false}

	_ = database.CreateNewToDoItem(todoItem)
}

func setItemComplete(itemId int) {
	item := database.GetItem(itemId)
	item.IsComplete = true

	database.UpdateToDoItem(item)
}
