package todo

import (
	"todo/ent"
	"todo/internal/database"
)

type Todo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// CreateTodo create todo in the database
func CreateTodo(t Todo) (*ent.Todo, error) {
	ct, err := database.DBClient.Todo.Create().
		SetTask(t.Task).
		Save(database.DBCtx)
	if err != nil {
		return nil, err
	}

	return ct, err
}
