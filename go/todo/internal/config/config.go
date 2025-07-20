package config

import (
  "log"
  "github.com/joho/godotenv"
  "os"
)
type Config struct {
  Port string
  DBUrl string
}
// godotenv.Load() returns nil on success and error on failure
func LoadConfig() *Config {
  err := godotenv.Load()
  if err != nil {
  log.Println("Error Loading Dot env")
}
  cfg := &Config{
  Port: GetEnv("PORT"),
  DBUrl: GetEnv("DB_URL"),
}
  return cfg
}
//this willl be a helper function to get ENV
func GetEnv(key string) string {
  val,status := os.LookupEnv(key)
  if !status {
    log.Println("NO ENV FOUND")
    return ""
  }
  return val
}


