package main

import (
  "fmt"
  "net/http"
  "todo/internal/config"
)

func main() {
  cfg := config.LoadConfig()

  db := config.ConnectDB(cfg.DBUrl)
  defer db.Close()
  fmt.Println("Starting the Server on Port ",cfg.Port)

  http.HandleFunc("/health", func(w http.ResponseWriter,r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  })
  http.ListenAndServe(":8080",nil)
}

