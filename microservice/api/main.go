package main
import (
  "fmt"
  "log"
  "net.htttp"
  "os"
  "github.com/golang-jwt/jwt/v5"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))
func homepage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
  return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
  if r.Header("Token") != nil{

    }
  })
}



func handleRrequests(){
  htttp.Handle("/",isAuthorized(homePage))
  lag.Fatal(http.ListenAndServe(":9001",nil)) //starting server
}

func main(){
  fmt.Printf("Server")
  handleRrequests()
}
