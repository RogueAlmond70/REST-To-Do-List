package Model

import (
	"fmt"
	"sort"
	"time"
)

type ToDoItem struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

type Request struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Completed   interface{} `json:"completed"`
}

type Database struct {
	items []ToDoItem
}

func (D *Database) Len() int {
	return len(D.items)
}

func (D *Database) Less(i, j int) bool {
	return D.items[i].CreatedAt.Before(D.items[j].CreatedAt)
}

func (D *Database) Swap(i, j int) {
	D.items[i], D.items[j] = D.items[j], D.items[i]
}

var instance *Database

func GetDatabase() *Database {
	if instance == nil {
		instance = &Database{}
	}
	return instance
}

func GetToDos(D *Database) *Database {
	sort.Sort(D)
	return D
}

// The http handler will return a 404 error when this function returns an error
func GetToDoByID(D *Database, ID int) (ToDoItem, error) {
	var Data = GetDatabase()
	for _, i := range Data.items {
		if ID == i.ID {
			return i, nil
		}
	}
	return ToDoItem{}, fmt.Errorf("item not found")
}
