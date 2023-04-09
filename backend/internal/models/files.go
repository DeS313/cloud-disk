package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Files struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Type       string             `bson:"type"`
	AccessLink string             `bson:"access_link"`
	Path       string             `bson:"path"`
	Size       int                `bson:"size"`
	UserID     primitive.ObjectID `bson:"userID"`
	ParrentID  primitive.ObjectID `bson:"parrentID"`
}

type CreateFile struct {
	Name   string             `json:"name"`
	Type   string             `json:"type"`
	Parent primitive.ObjectID `json:"parent"`
}
