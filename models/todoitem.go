package models

type ToDoItem struct {
	Id          int
	Name        string
	Description string
	IsComplete  bool
}

func NewToDoItem(id int, name string, description string) *ToDoItem {
	item := ToDoItem{Id: id, Name: name, Description: description, IsComplete: false}
	return &item
}
