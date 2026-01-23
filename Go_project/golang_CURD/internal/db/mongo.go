package db

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {
	//prevent your app from freezing in startup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOps := options.Client().ApplyURI(cfg.MongoUrl)

	client, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		return nil, nil, fmt.Errorf("mongo connect failed")
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("mongo ping failed")
	}
	database := client.Database(cfg.MongoDB)

	return client, database, nil

}

func Disconnet(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}
