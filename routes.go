package main

import (
	"log"
	"net/http"

	handlers "webscoket/handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Here we will register application routes. Here we will use gorilla/mux package to register routes

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public").Handler(http.StripPrefix("/public/", fs))
}

// AddApproutes will add the routes fro the eapplication
func AddApproutes(route *mux.Router) {
	log.Println("Loading routes...")
	setStaticFolder(route)
	hub := handlers.NewHub()
	go hub.Run()

	route.HandleFunc("/", handlers.RenderHome)
	route.HandleFunc("/ws/{username}", func(responseWriter http.ResponseWriter, request *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		username := mux.Vars(request)["username"]
		connection, err := upgrader.Upgrade(responseWriter, request, nil)
		if err != nil {
			log.Println("Error occured::", err)
			return
		}
		handlers.CreateNewSocketUser(hub, connection, username)
	})
	log.Println("Routes are loaded...")
}
