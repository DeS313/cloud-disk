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

func (s *Storage) FindOne(ctx context.Context, id string) (models.User, error) {
	var user models.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, fmt.Errorf("ошибка преобразования hex в ObjectID, hex: %s", id)
	}

	result := s.db.FindOne(ctx, bson.M{"_id": oid})
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

func (s *Storage) Create(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	log.Println("созданиие пользователя")
	_, err := s.db.InsertOne(ctx, user)

	if err != nil {
		return user.ID, fmt.Errorf("ошибка создания пользователя: %v", err)
	}

	return user.ID, nil

}
