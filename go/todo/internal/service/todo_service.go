package service 

import (
  "context"
  "time"
  "todo/internal/model"
  "todo/internal/repositry"
)

type TodoService interface {
    CreateTodo(ctx context.Context, title string) (*model.Todo, error)
    GetAllTodo(ctc context.Context) ([]*model.Todo,error)
}
type todoService struct {                     //intialise to the repo as
    repo repositry.TodoRepository
}

func NewTodoService(repo repositry.TodoRepository) TodoService {  //returns a new service of type TodoService
    return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(ctx context.Context, title string) (*model.Todo, error) {
    todo := &model.Todo{
        Title:     title,
        Completed: false,
        CreatedAt: time.Now(),
    }

    err := s.repo.CreateTodo(ctx, todo)
    return todo, err

func (s *todoService) GetAllTodo(ctx context.Context) ([]*model.Todo,error)
  {
    err := s.repo.GetAllTodos(ctx)
  }
//FUll WORKFLOW
/*main.go
  |
  v
HTTP handler (uses TodoService)
  |
  v
TodoService interface
  |
  v
todoService struct (with repo injected)
  |
  v
TodoRepository interface
  |
  v
todoRepo struct (with db access)
}*/
}
