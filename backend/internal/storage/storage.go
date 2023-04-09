package storage

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database, collection string) *Storage {
	return &Storage{
		db: db,
	}
}
