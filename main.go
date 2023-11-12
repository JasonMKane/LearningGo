package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"todo/database"
	"todo/models"

	"github.com/jedib0t/go-pretty/table"
)

var (
	list     = flag.Bool("list", false, "List all items in To-Do")
	new      = flag.Bool("new", false, "Create a new to-do items")
	complete = flag.Int("complete", 0, "Id for the item to mark complete")
	show     = flag.Int("show", -1, "Id for the item to show")
)

func main() {
	// todo := database.GetItem(1)
	// fmt.Println(todo)
	flag.Parse()

	cmdArgs := os.Args[1:]
	if len(cmdArgs) == 0 {
		fmt.Println(`
			No argument specified.

			Options:
			
			list: show all To-Do Items

			new: Create a new Item

			complete: mark a to-do item as complete.
		`)
	}

	if isCommand("list", cmdArgs) {
		listAllOpenItems()
	}

	if isCommand("new", cmdArgs) {
		createNewItem()
	}

	if isCommand("-show", cmdArgs) {
		showItem(cmdArgs)
	}

	if isCommand("complete", cmdArgs) {
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

func isCommand(cmdName string, args []string) bool {
	if slices.Contains(args, cmdName) {
		return true
	} else {
		return false
	}
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

func showItem(cmdArgs []string) {
	item := database.GetItem(*show)
	items := [1]models.ToDoItem{*item}
	printToDoItem(items)
}

func printToDoItem(items []models.ToDoItem) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Description", "IsComplete"})
	// t.AppendSeparator()
	for _, v := range items {
		t.AppendRow([]interface{}{v.Id, v.Name, v.Description, v.IsComplete})
	}

	t.Render()

	// fmt.Printf("%v  |", item.Id)
	// fmt.Printf(" %v  |", item.Name)
	// fmt.Printf(" %v  |", item.Description)
	// fmt.Printf(" %v \n", item.IsComplete)
}
