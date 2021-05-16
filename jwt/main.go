package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Create this on an environment variable
// var mySigningKey = os.Get("JWT_TOKEN")
var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Carlos Nava"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong %s", err.Error())
		return "", err
	}

	return tokenString, err
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("My simple client")

	handleRequests()
}
