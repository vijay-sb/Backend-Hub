
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY")) // FIXED []byte[...] to []byte(...)

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256) // FIXED: SigningMethodHS526 → SigningMethodHS256

	claims := token.Claims.(jwt.MapClaims)

	claims["authorised"] = true
	claims["client"] = "Vijay SB"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("Something went wrong: %s", err.Error()) // FIXED: added return
	}
	return tokenString, nil // FIXED: added nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to Generate the Token")
	}
	fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", Index) // FIXED: htttp → http, index → Index (exported)
	log.Fatal(http.ListenAndServe(":8080", nil)) // FIXED: htttp → http
}

func main() { // FIXED: main() block syntax
	handleRequests()
}
