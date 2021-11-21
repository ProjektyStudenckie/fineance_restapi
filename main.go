package main

import (
	http2 "ApiRest/internal/http"
	"ApiRest/internal/mongo"
	"fmt"
	"net/http"
)


type App struct{
}


func (app *App) Run() error {
	fmt.Println("Setting Up Rest Api")
	mongo.DataBaseCon.SetupDataBaseConnection()
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

}


