package storage

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	db *mongo.Collection
}

func NewStorage(db *mongo.Database, collection string) *Storage {
	return &Storage{
		db: db.Collection(collection),
	}
}
