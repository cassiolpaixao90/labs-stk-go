package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(m.URI)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			return
		}

		mongoClient = client
	})

	return mongoClient
}
