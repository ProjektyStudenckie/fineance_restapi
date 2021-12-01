package Wallet

import (
	"ApiRest/internal/mongo"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type RequestStruct struct {
	Wallet Wallet `json:"wallet,omitempty" bson:"wallet ,omitempty"`
	Owner mongo.User  `json:"owner,omitempty" bson:"owner ,omitempty"`
}

type Wallet struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string  				`json:"name,omitempty" bson:"name ,omitempty"`
	Owner mongo.User  				`json:"owner,omitempty" bson:"owner ,omitempty"`
	Currency string             `json:"currency,omitempty" bson:"currency ,omitempty"`
	SubOwners  []mongo.User            `json:"subowners,omitempty" bson:"subowners,omitempty"`
	Value  int             `json:"value,omitempty" bson:"value,omitempty"`
}


func AddWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var wallet Wallet

	_ = json.NewDecoder(request.Body).Decode(&wallet)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection.InsertOne(ctx,wallet)
}

func AddSubOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var reqstruct RequestStruct
	_ = json.NewDecoder(request.Body).Decode(&reqstruct)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cur := collection.FindOne(ctx, Wallet{ID: reqstruct.Wallet.ID})
	json.NewEncoder(response).Encode(cur)
}


func GetWallets(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var wallets []Wallet

	var user mongo.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cur,err := collection.Find(ctx, Wallet{Owner: mongo.User{ Username:user.Username}})

	if err!=nil{
		fmt.Println(err)
		return
	}
	err = cur.All(ctx,wallets)
	if err!=nil{
		fmt.Println(err)
		return
	}

	json.NewEncoder(response).Encode(wallets)
	return
}


func GetSubWallets(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var wallets []Wallet

	var user mongo.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cur,err := collection.Find(ctx, Wallet{SubOwners: []mongo.User{{Username:user.Username }}})

	if err!=nil{
		fmt.Println(err)
		return
	}
	err = cur.All(ctx,wallets)
	if err!=nil{
		fmt.Println(err)
		return
	}

	json.NewEncoder(response).Encode(wallets)
	return
}