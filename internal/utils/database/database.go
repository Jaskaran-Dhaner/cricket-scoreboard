package database

import (
	"scoreboard/config"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(collectionName string) *mongo.Collection {
	collection := config.AppConfig.Database.DB.Collection(collectionName)
	return collection
}
