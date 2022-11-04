package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opchav/todos-go-api/app/models"
)

// TodoQueries struct for queries from Todo model
type TodoQueries struct {
	*sqlx.DB
}

// GetTodos method for getting all todos.
func (q *TodoQueries) GetTodos() ([]models.Todo, error) {
	todos := []models.Todo{}

	query := `SELECT * FROM todos`

	err := q.Select(&todos, query)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

// GetTodo method for getting one todo by given ID.
func (q *TodoQueries) GetTodo(id uuid.UUID) (models.Todo, error) {
	todo := models.Todo{}

	query := `SELECT * FROM todos WHERE id = $1`

	err := q.Get(&todo, query, id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

// CreateTodo method for creating todo by given Todo object.
func (q *TodoQueries) CreateTodo(t *models.Todo) error {
	query := `INSERT INTO todos VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := q.Exec(query, t.ID, t.CreatedAt, t.UpdatedAt, t.UserID, t.Content, t.Complete, t.TodoAttrs)
	if err != nil {
		return err
	}

	return nil
}

// UpdatedTodo method for updating todo by given Todo object.
func (q *TodoQueries) UpdatedTodo(id uuid.UUID, t *models.Todo) error {
	query := `UPDATE todos SET updated_at = $2, content = $3, complete = $4, todo_attrs = $5 WHERE id = $1`

	_, err := q.Exec(query, id, t.UpdatedAt, t.Content, t.Complete, t.TodoAttrs)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodo method for deleting todo by given ID.
func (q *TodoQueries) DeleteTodo(id uuid.UUID) error {
	query := `DELETE FROM todos WHERE id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
