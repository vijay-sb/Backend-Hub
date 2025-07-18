package main
import (
  "fmt"
  "log"
  "net.htttp"
  "os"
  jwt "github.com/golang-jwt/jwt/v5"
)
var mySigningKey = []byte[os.Getenv("SECRET_KEY")]

func GetJWT()(string,error){
  token := jwt.New(jwt.SigningMethodHS526)
  claims := token.Claims.(jwt.MapClaims)

  
  claims["authorised"] = true
  claims["client"] = "Vijay SB"
  claims["aud"] = "billing.jwtgo.io"
  claims["iss"] = "jwtgo.io"
  claims["exp"] = time.Now().Add(time.Minute*1).Unix()

  tokenString, err := token.SignedString(mySigningKey)

  if err != nil{
    fmt.Errorf("Something went wrong: %s",err.Error())
    return "", err
  }
  return tokenString, nil

}

func Index(w http.ResponseWriter,r *http.Request){

  validToken, err := GetJWT()
  fmt.Println(validToken)
  if err != nil{
    fmt.Println("Failed to Genrate the Token")
  }
  fmt.Fprintf(w,string(validToken))
  
}

func handleRequests(){
htttp.HandleFunc("/",index)

log.
}
