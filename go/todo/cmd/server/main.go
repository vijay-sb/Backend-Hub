package main

import (
      "database/sql"
    "log"
  "fmt"
  "net/http"
      "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
  "todo/internal/config"
  "todo/internal/handler"
  "todo/internal/repositry"
  "todo/internal/service"
)

func main() {
  cfg := config.LoadConfig()

  db := config.ConnectDB(cfg.DBUrl)
  defer db.Close()
  fmt.Println("Starting the Server on Port ",cfg.Port)
  todoRepo := repositry.NewTodoRepository(db)
  todoSvc := service.NewTodoService(todoRepo)
  todoHandler := handler.NewServiceHandler(todoSvc)

  r := chi.NewRouter()
  r.Use(middleware.Logger)

  r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

  r.Route("/todos", func(r chi.Router) {
        r.Post("/", todoHandler.CreateTodo)
        r.Get("/", todoHandler.GetAllTodos)
    })
  if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }


}

