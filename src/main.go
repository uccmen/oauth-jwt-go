package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login-callback", loginCallbackHandler)
	http.HandleFunc("/me", ValidateMiddleware(meHandler))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
