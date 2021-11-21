package main

import (
	http2 "ApiRest/internal/http"
	mongo2 "ApiRest/internal/mongo"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)


type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Rest Api")
	handler := http2.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":1332", handler.Router); err !=nil{
		fmt.Println("Failed to setup")
		return err
	}
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

}


