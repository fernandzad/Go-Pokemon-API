package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "root"
	pwd      = "Inception2575.-"
	database = "pokemon"
)

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@testgo.6itcv.mongodb.net/%s?retryWrites=true&w=majority", usr, pwd, database)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer client.Disconnect(ctx)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Databases: ", databases)

	return client.Database(database).Collection(collection)
}
