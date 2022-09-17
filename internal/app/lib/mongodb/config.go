package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const Database = "rio_bus"

func New() *mongo.Database {
	URL := os.Getenv("MONGO_URL")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))
	if err != nil {
		panic(err)
	}

	return client.Database(Database)
}
