package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var client *mongo.Client


func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("cluster0-shard-00-02.a3zgx.mongodb.net:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Starting the NewRouter...")
	router := mux.NewRouter()
	router.HandleFunc("/person", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/person/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/login" ,Login).Methods("POST")
	http.ListenAndServe(":1332", router)

}