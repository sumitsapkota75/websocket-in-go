package main

import (
	"fmt"
	"log"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page for websockets")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func SetupRoutes() {
	http.HandleFunc("/", Homepage)
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// remove the previous fmt statement
	// fmt.Fprintf(w, "Hello World")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

}

func main() {
	fmt.Println("Websockets initialized")
	SetupRoutes()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
