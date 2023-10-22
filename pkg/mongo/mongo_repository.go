package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoRepository[I any, O any] interface {
	InsertOne(schema *I, ctx context.Context) (*O, error)
}

type Repository[I any, O any] struct {
	collection *mongo.Collection
}

func NewRepository[I any, O any](client *mongo.Client, dbName, collectionName string) IMongoRepository[I, O] {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repository[I, O]{collection: collection}
}

func (r *Repository[I, O]) InsertOne(schema *I, ctx context.Context) (*O, error) {
	_, err := r.collection.InsertOne(ctx, schema)
	return nil, err
}

func (r *Repository[I, O]) FindAll(schema *[]I, ctx context.Context) (*[]O, error) {

	_, err := r.collection.Find(ctx, schema)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (r *Repository[I, O]) UpdateOne(schema *I, filter interface{}, ctx context.Context) (*O, error) {
	_, err := r.collection.UpdateOne(ctx, filter, schema)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
