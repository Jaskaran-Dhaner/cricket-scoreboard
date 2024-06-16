package database

import (
	"context"
	"scoreboard/config"
	"scoreboard/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func init() {
	config.AppConfig.Database.DB, _ = LoadDB()
	logger.Info("database loaded")
}

func getClient() (*mongo.Client, error) {
	var err error
	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(config.Database.MongoURI)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			logger.Fatal(err.Error())
		}

		// Ping the database to verify connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Ping(ctx, nil); err != nil {
			logger.Fatal(err.Error())
		}
		clientInstance = client
	})
	return clientInstance, err
}

func LoadDB() (*mongo.Database, error) {
	client, err := getClient()
	if err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}

	config.AppConfig.Database.Client = client
	// Set client options
	logger.Info("connected to mongodb")
	return config.AppConfig.Database.Client.Database(config.Database.DBName), nil
}
