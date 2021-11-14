package main

import (
	http2 "ApiRest/internal/http"
	mongo2 "ApiRest/internal/mongo"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)


type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Rest Api")
	return nil
}

func main() {
	fmt.Println("Starting the application...")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting Rest Api")
		fmt.Println(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://Wielok:Projekt123@cluster0.a3zgx.mongodb.net/TestDB?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongo2.Client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Starting the NewRouter...")
	router := mux.NewRouter()
	router.HandleFunc("/user", mongo2.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id}", mongo2.GetUserEndpoint).Methods("GET")
	router.HandleFunc("/login/{password}/{username}", http2.Login).Methods("POST")
	router.HandleFunc("/test/{test}", http2.Test).Methods("GET")
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

