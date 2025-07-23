package main

import (
  "fmt"
  "net/http"
  "todo/internal/config"
  "todo/internal/handler"
  "todo/internal/repository"
  "todo/internal/service"
)

func main() {
  cfg := config.LoadConfig()

  db := config.ConnectDB(cfg.DBUrl)
  defer db.Close()
  fmt.Println("Starting the Server on Port ",cfg.Port)
  todoRepo := repository.NewTodoRepository(db)
  todoSvc := service.NewTodoService(todoRepo)
  todoHandler := handler.NewTodoHandler(todoSvc)

  http.HandleFunc("/health", func(w http.ResponseWriter,r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  })

  http.HandleFunc("/todos", todoHandler.CreateTodo)
  http.ListenAndServe(":8080",nil)
}

