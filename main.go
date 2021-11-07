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
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/login/{password}/{username}" ,Login).Methods("POST")
	router.HandleFunc("/test/{test}" ,Test).Methods("GET")
	http.ListenAndServe(":1332", router)

}