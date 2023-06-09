package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" `
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	DiskSpace int                `bson:"diskSpace"`
	UserSpace int                `bson:"userSpace"`
	Avatar    string             `bson:"avatar"`
	Files     []Files
}

type CreateUser struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
