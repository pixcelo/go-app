package database

import (
	"context"
	"fmt"
	"go/mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"

	"github.com/joho/godoenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBinstance() *mongo.Client {
	err := godoenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 19*time.Second)
	defer cancel()
	err = client.Cotext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB.")

	return client
}

// MongoDB Instance
var Client *mongo.client = DBinstance()

func OpenCollection(client *mongo.client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.database("cluster0").Collection(collectionName)
	return collection
}
