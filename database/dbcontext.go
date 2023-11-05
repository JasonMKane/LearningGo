package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"todo/models"

	"github.com/jackc/pgx/v5"
)

func getConnection() *pgx.Conn {
	dbUrl := "postgres://db_user:db_password@localhost:5432/ToDo"
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer conn.Close(context.Background())

	return conn
}

func GetItem(itemId int) *models.ToDoItem {

	conn := getConnection()

	var id int
	var name string
	var description string
	err := conn.QueryRow(context.Background(), "SELECT \"Id\", \"Name\", \"Description\" FROM \"Items\" where \"Id\"=$1", itemId).Scan(&id, &name, &description)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetItem failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(id)
	defer conn.Close(context.Background())
	return models.NewToDoItem(id, name, description)
}

func GetOpenToDoItems() []models.ToDoItem {
	conn := getConnection()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM \"Items\" WHERE \"IsComplete\" = FALSE")
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateNewToDoItem failed: %v\n", err)
	}

	defer rows.Close()

	var rowSlice []models.ToDoItem

	for rows.Next() {
		var r models.ToDoItem
		err := rows.Scan(&r.Id, &r.Name, &r.Description, &r.IsComplete)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return rowSlice
}

func CreateNewToDoItem(item models.ToDoItem) *models.ToDoItem {
	conn := getConnection()
	var id int
	err := conn.QueryRow(context.Background(), "INSERT INTO \"Items\" (\"Name\", \"Description\") VALUES ($1, $2) RETURNING \"Id\"", item.Name, item.Description).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateNewToDoItem failed: %v\n", err)
	}
	defer conn.Close(context.Background())

	return GetItem(id)
}

func UpdateToDoItem(item *models.ToDoItem) *models.ToDoItem {
	conn := getConnection()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "UPDATE \"Items\" SET \"Name\"=$1, \"Description\"=$2 WHERE \"Id\"=$3", item.Name, item.Description, item.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "UpdateToDoItem failed: %v\n", err)
	}

	return GetItem(item.Id)
}
