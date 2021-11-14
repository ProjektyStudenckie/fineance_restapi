package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var client *mongo.Client

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://Wielok:Projekt123@cluster0.a3zgx.mongodb.net/TestDB?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Starting the NewRouter...")
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/login/{password}/{username}", Login).Methods("POST")
	router.HandleFunc("/test/{test}", Test).Methods("GET")
	http.ListenAndServe(":1332", router)

}

func Writer(conn *websocket.Conn) {

	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Println(t)
		}
	}
}
