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
      jwt.Parse(r.Header["Token"][0],func(token *jwt.Token)(interface{},error){
        if _,ok := token.,Method(*jwt.SigningMethodHMAC); !ok{
          return nil,fmt.Errorof(("Invalid signing Method"))
        }
        aud := "billing.jwtgo.io"
        
        checkAudience:=token.CLaims.(jwt.MapClaims).VerifyAudience(aud,false)
        if !checkAudience {
          return nil, fmt.Errorof(("Invalid aud"))

        }
        iss:="jwtgo.io"
        
        checkIss := token.Claims.(jwt.MapClaims)).VerifyIssuer(iss, false)
        if !checkIss{
          return nil,fmt.Errorof(("Invalid iss"))
        }

      })
      if err != nil{
        fmt.Printf(w,err.Rrror())
      }
      if token.valid{
        endpoint(w,r)

      }

    }
          else{
        fmt.Printf(w,"No Auth Token Provided")
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
