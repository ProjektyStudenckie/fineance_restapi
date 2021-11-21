package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DataBaseCon DataBaseConnection

type DataBaseConnection struct{
	Client *mongo.Client
}



type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username ,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	RT  string             `json:"rt,omitempty" bson:"rt,omitempty"`
}


func (h *DataBaseConnection) SetupDataBaseConnection(){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://Wielok:Projekt123@cluster0.a3zgx.mongodb.net/TestDB?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	h.Client, _ = mongo.Connect(ctx, clientOptions)
}