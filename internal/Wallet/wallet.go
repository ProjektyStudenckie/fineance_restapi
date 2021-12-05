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
	Wallet Wallet     `json:"wallet,omitempty" bson:"wallet ,omitempty"`
	User   mongo.User `json:"user,omitempty" bson:"user ,omitempty"`
}

type RequestStructWalletGoal struct {
	Wallet Wallet     `json:"wallet,omitempty" bson:"wallet ,omitempty"`
	Goals   Goals `json:"goal,omitempty" bson:"goal ,omitempty"`
}

type RequestStructWalletRemittance struct {
	Wallet Wallet     `json:"wallet,omitempty" bson:"wallet ,omitempty"`
	Remittance   Remittance `json:"remittance,omitempty" bson:"remittance ,omitempty"`
}

type Wallet struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string  				`json:"name,omitempty" bson:"name ,omitempty"`
	Description string `json:"description,omitempty" bson:"description ,omitempty"`
	Owner mongo.User  				`json:"owner,omitempty" bson:"owner ,omitempty"`
	Currency string             `json:"currency,omitempty" bson:"currency ,omitempty"`
	SubOwners  []mongo.User            `json:"subowners,omitempty" bson:"subowners,omitempty"`
	WalletGoals []Goals		`json:"goals,omitempty" bson:"goals,omitempty"`
	Value  []Remittance             `json:"remittance,omitempty" bson:"remittance,omitempty"`
}

type Goals struct{
	Name string  				`json:"name,omitempty" bson:"name ,omitempty"`
	Date string 			`json:"date,omitempty" bson:"date ,omitempty"`
	Value  int             `json:"value,omitempty" bson:"value,omitempty"`
}

type Remittance struct{
	Date string 			`json:"date,omitempty" bson:"date ,omitempty"`
	Value  int             `json:"value,omitempty" bson:"value,omitempty"`
}


func AddWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var wallet Wallet

	_ = json.NewDecoder(request.Body).Decode(&wallet)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection.InsertOne(ctx,wallet)
	json.NewEncoder(response).Encode(true)
}

func AddSubOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var reqstruct RequestStruct
	var wallet Wallet
	_ = json.NewDecoder(request.Body).Decode(&reqstruct)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_= collection.FindOne(ctx, Wallet{ID: reqstruct.Wallet.ID}).Decode(&wallet)
	wallet.SubOwners=append(wallet.SubOwners, reqstruct.User)

	_,_ =collection.ReplaceOne(ctx,Wallet{ID: wallet.ID},
		wallet,
	)
	json.NewEncoder(response).Encode(true)
}

func AddGoal(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var reqstruct RequestStructWalletGoal
	var wallet Wallet
	_ = json.NewDecoder(request.Body).Decode(&reqstruct)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_= collection.FindOne(ctx, Wallet{ID: reqstruct.Wallet.ID}).Decode(&wallet)
	wallet.WalletGoals=append(wallet.WalletGoals, reqstruct.Goals)

	_,_ =collection.ReplaceOne(ctx,Wallet{ID: wallet.ID},
		wallet,
	)
	json.NewEncoder(response).Encode(true)
}

func RemoveGoal(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var reqstruct RequestStructWalletGoal
	var wallet Wallet
	_ = json.NewDecoder(request.Body).Decode(&reqstruct)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_= collection.FindOne(ctx, Wallet{ID: reqstruct.Wallet.ID}).Decode(&wallet)
	index:= posGoal(reqstruct.Goals,reqstruct.Wallet.WalletGoals)
	wallet.WalletGoals = removeGoal(wallet.WalletGoals,index)
	_,_ =collection.ReplaceOne(ctx,Wallet{ID: wallet.ID},
		wallet,
	)
	json.NewEncoder(response).Encode(true)
}


func AddRemittance(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var reqstruct RequestStructWalletRemittance
	var wallet Wallet
	_ = json.NewDecoder(request.Body).Decode(&reqstruct)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_= collection.FindOne(ctx, Wallet{ID: reqstruct.Wallet.ID}).Decode(&wallet)
	wallet.Value=append(wallet.Value, reqstruct.Remittance)

	_,_ =collection.ReplaceOne(ctx,Wallet{ID: wallet.ID},
		wallet,
	)
	json.NewEncoder(response).Encode(true)
}

func UpdateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var wallet Wallet
	_ = json.NewDecoder(request.Body).Decode(&wallet)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_,err :=collection.ReplaceOne(ctx,Wallet{ID: wallet.ID},
		wallet,
	)
	if err!=nil {
		json.NewEncoder(response).Encode(false)
		return
	}
	json.NewEncoder(response).Encode(true)
}



func RemoveSubOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var reqstruct RequestStruct
	var wallet Wallet
	_ = json.NewDecoder(request.Body).Decode(&reqstruct)
	collection := mongo.DataBaseCon.Client.Database("TestDB").Collection("Wallets")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_= collection.FindOne(ctx, Wallet{ID: reqstruct.Wallet.ID}).Decode(&wallet)
	index:= pos(reqstruct.User,reqstruct.Wallet.SubOwners)
	wallet.SubOwners = remove(wallet.SubOwners,index)
	_,_ =collection.ReplaceOne(ctx,Wallet{ID: wallet.ID},
		wallet,
	)
	json.NewEncoder(response).Encode(true)
}




func remove(s []mongo.User, i int) []mongo.User {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func removeGoal(s []Goals, i int) []Goals {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func pos(user mongo.User, value []mongo.User) int {
	for p, v := range value {
		if v == user {
			return p
		}
	}
	return -1
}

func posGoal(goals Goals, value []Goals) int {
	for p, v := range value {
		if v == goals {
			return p
		}
	}
	return -1
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
	err = cur.All(ctx,&wallets)
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
	if len(wallets)>0{
		json.NewEncoder(response).Encode(&wallets)
	} else{
		json.NewEncoder(response).Encode("no wallets")
		return
	}

	return
}