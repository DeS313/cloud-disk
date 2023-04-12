package storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/DeS313/cloud-disk/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const users = "users"

func (s *Storage) FindOne(ctx context.Context, id string) (models.User, error) {
	var user models.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, fmt.Errorf("ошибка преобразования hex в ObjectID, hex: %s", id)
	}

	result := s.db.Collection(users).FindOne(ctx, bson.M{"_id": oid})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&user); err != nil {
		return models.User{}, fmt.Errorf("failed to decode user (id: %s) from DB due to error: %v", id, err)
	}

	return user, nil

}

func (s *Storage) FindOneByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	result := s.db.Collection(users).FindOne(ctx, bson.M{"email": email})

	if err := result.Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (s *Storage) Create(ctx context.Context, user models.User) (string, error) {
	log.Println("созданиие пользователя")
	result, err := s.db.Collection(users).InsertOne(ctx, user)

	if err != nil {
		log.Println(fmt.Errorf("ошибка создания пользователя: %v", err))
		return user.ID.Hex(), err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert ObjectID")
	}

	return oid.Hex(), nil
}

func (s *Storage) Update(ctx context.Context, user *models.User) error {
	fmt.Println("user", user)
	filter := bson.M{"_id": user.ID}

	update := bson.M{
		"$set": bson.M{
			"userSpace": user.UserSpace,
		},
	}

	result, err := s.db.Collection(users).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("ERR: result matched count")
	}

	return nil
}
