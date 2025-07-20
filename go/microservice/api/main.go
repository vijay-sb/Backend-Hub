
package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "github.com/golang-jwt/jwt/v5"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homepage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Token") != "" {
      token, err := jwt.Parse(r.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("Invalid signing Method")
        }

        claims := token.Claims.(jwt.MapClaims)

        aud := "billing.jwtgo.io"
        if claims["aud"] != aud {
          return nil, fmt.Errorf("Invalid aud")
        }

        iss := "jwtgo.io"
        if claims["iss"] != iss {
          return nil, fmt.Errorf("Invalid iss")
        }

        return MySigningKey, nil
      })

      if err != nil {
        fmt.Fprintf(w, err.Error())
      }

      if token != nil && token.Valid {
        endpoint(w, r)
      }

    } else {
      fmt.Fprintf(w, "No Auth Token Provided")
    }
  })
}

func handleRrequests() {
  http.Handle("/", isAuthorized(homepage))
  log.Fatal(http.ListenAndServe(":9001", nil)) // starting server
}

func main() {
  fmt.Printf("Server")
  handleRrequests()
}
