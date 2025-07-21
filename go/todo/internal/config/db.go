package db

import (
   "database/sql"
    "log"
    "time"

    _ "github.com/jackc/pgx/v5/stdlib"
)
func ConnectDB(dbURL string) *sql.DB{
  //return *sql.DB pointer to db
  db,err := sql.Open("pgx",dbURL)
  if err != nil {
    log.Fatalf("Failed to open %v",err);
  }
   db.SetMaxOpenConns(10)                 //Ten connections get opened in pool
   db.SetMaxIdleConns(5)                  //5 out of 10 is kept alive 
  db.SetConnMaxLifetime(time.Hour)       //after one hour close a connection and restart again
  if err := db.Ping(); err != nil {
    log.Fatalf("Failed to ping DB: %v", err)
    }
  log.Println("Connection to pg is Sucessfull")
  return db
}

