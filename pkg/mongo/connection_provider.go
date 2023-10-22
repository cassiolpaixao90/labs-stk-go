package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

type MongoConfig struct {
	URI    string
	Logger interface{}
}

func NewMongoConnection(cfg *MongoConfig) *MongoConfig {
	return &MongoConfig{URI: cfg.URI}
}

func (m *MongoConfig) GetMongoClient() *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(m.URI)

		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			return
		}

		mongoClient = client
	})

	return mongoClient
}
