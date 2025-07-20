package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("Starting the Server on Port: 8080")

  http.HandleFunc("/health", func(w http.ResponseWriter,r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  })
  http.ListenAndServe(":8080",nil)
}

