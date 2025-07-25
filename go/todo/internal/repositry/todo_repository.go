// using repositry desing pattern
package repositry

import (
	"context"
	"database/sql"
	"todo/internal/model"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
}

type todoRepo struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepo{db: db}
}

func (r *todoRepo) CreateTodo(ctx context.Context, todo *model.Todo) error {
	query := `INSERT INTO todos (title, completed, created_at) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, todo.Title, todo.Completed, todo.CreatedAt).Scan(&todo.ID)
	return err
}

func (r *todoRepo) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id,title,completed,created_at FROM todos ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

// todoRepo needs to have all funcitons of interface TodoRepository
// and (r *todoRepo)  this line make the CreateTodo funciton releate to the interface funciton
// When using this first use the NewTodoRepository fucntion to intialise db value then use funciton from the interface
