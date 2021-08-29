package db

import (
	"context"
	"fmt"
	"time"

	"github.com/MojixCoder/awesomeProject/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDb connects to mongodb and returns client
func GetDb() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetConfig().MongoURI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// GetCollection returns a db collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(config.GetConfig().DBName).Collection(collectionName)
	return collection
}
