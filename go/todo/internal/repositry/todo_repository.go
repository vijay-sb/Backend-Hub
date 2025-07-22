// using repositry desing pattern
package repositry
import (
    "context"
    "database/sql"
    "todo/internal/model"
)

type TodoRepository interface {
  CreateTodo(ctx context.Context,todo *model.Todo) error
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
