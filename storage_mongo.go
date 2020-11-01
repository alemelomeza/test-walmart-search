package main

import (
	"context"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// CollectionProducts identifier for collection about products
	CollectionProducts = "products"
)

// StorageMongo ...
type StorageMongo struct {
	db *mongo.Client
}

// NewStorageMongo ...
func NewStorageMongo(uri string) (*StorageMongo, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	stg := new(StorageMongo)
	stg.db = client
	return stg, nil
}

// Find ...
func (s *StorageMongo) Find(criteria string) ([]*Product, error) {
	id, err := strconv.Atoi(criteria)
	if err != nil {
		return s.findByText(criteria)
	}

	return s.findByID(id)
}

// findByID ...
func (s *StorageMongo) findByID(criteria int) ([]*Product, error) {
	collection := s.db.Database("promotions").Collection(CollectionProducts)
	var ps []*Product
	var p Product

	filter := bson.M{"id": criteria}
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		return ps, err
	}
	ps = append(ps, &p)

	return ps, nil
}

// findByText ...
func (s *StorageMongo) findByText(criteria string) ([]*Product, error) {
	collection := s.db.Database("promotions").Collection(CollectionProducts)
	var ps []*Product

	filter := bson.M{
		"$or": []bson.M{
			bson.M{"brand": primitive.Regex{Pattern: criteria, Options: "i"}},
			bson.M{"description": primitive.Regex{Pattern: criteria, Options: "i"}},
		},
	}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return ps, nil
	}
	for cur.Next(context.TODO()) {
		var p Product
		err := cur.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}
		ps = append(ps, &p)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return ps, nil
}
