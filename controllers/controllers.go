package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-mongo/connections"
	"go-mongo/models"
)

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		userCollection = connections.DB().Database("userDB").Collection("userClient")
		person         models.User
	)

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Print(err)
	}

	insertResult, err := userCollection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}

	// insertResult returns user_id
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		userCollection = connections.DB().Database("userDB").Collection("userClient")
		body           models.User
		result         primitive.M 
	)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Error while decoding, error: ", err.Error())
	}

	err = userCollection.FindOne(context.Background(), bson.D{{Key: "age", Value: body.Age}, {Key: "name", Value: body.Name}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		userCollection = connections.DB().Database("userDB").Collection("userClient")
		body           models.UpdateBody
	)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Print(err)
	}
	
	filter := bson.D{{Key: "name", Value: body.Name}}
	after := options.After

	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: body.Age}, {Key: "username", Value: body.Username}}}}
	updateResult := userCollection.FindOneAndUpdate(context.Background(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userCollection = connections.DB().Database("userDB").Collection("userClient")

	params := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(params)

	if err != nil {
		fmt.Println(err.Error())
	}

	opts := options.Delete().SetCollation(&options.Collation{})

	res, err := userCollection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(res.DeletedCount)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		userCollection = connections.DB().Database("userDB").Collection("userClient")
		results        []primitive.M
	)

	cur, err := userCollection.Find(context.Background(), bson.D{{}})
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.Background()) { 
		var elem primitive.M

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.Background()) 

	json.NewEncoder(w).Encode(results)
}
