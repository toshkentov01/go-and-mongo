package connections

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connnection to mongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("Error while connecting to mongodb, error: ", err.Error())
		panic(err)
	}
	
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("Error while checking connection to mongodb, error: ", err.Error())
		panic(err)
	}

	return client
}
