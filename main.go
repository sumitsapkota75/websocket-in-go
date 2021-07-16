package main

import (
	"fmt"
	"log"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page for websockets")
}

func SetupRoutes() {
	http.HandleFunc("/", Homepage)
}

func main() {
	fmt.Println("Websockets initialized")
	SetupRoutes()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
