package main
import (
  "fmt"
  "log"
  "net.htttp"
  "os"
  "github.com/golang-jwt/jwt/v5"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))
