package db

import (
	"context"
	"github.com/franciscof12/rest-api-thn/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"os"
)

func CreateMongoClient(uri string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		slog.Error("Error creating mongo client", "error", err.Error())
	}
	return client, nil
}

func GetMongoCollection(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func SetupDatabase() *mongo.Client {
	client, err := CreateMongoClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		slog.Error("Error creating mongo client:", "error", err.Error())
		os.Exit(1)
	}
	return client
}

func SetupModels(client *mongo.Client) {
	mongoCollection := GetMongoCollection(client, os.Getenv("DATABASE_NAME"), os.Getenv("COLLECTION_NAME"))
	models.SetRequestCollection(mongoCollection)
}
