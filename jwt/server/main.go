package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret info")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Validates if token header is set
		if r.Header["Token"] != nil {

			// Try to parse the token with HMAC enc algorithm and with signing key
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error.")
				}

				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			// If token is valid, render the homepage
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("My simple server")
	handleRequests()
}
