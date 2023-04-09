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

const cloudDisk = "cloud-disk"

func (s *Storage) CreateF(ctx context.Context, file models.Files) (string, error) {
	log.Println("созданиие файла")

	result, err := s.db.Collection(cloudDisk).InsertOne(ctx, file)

	if err != nil {
		log.Println(fmt.Errorf("ошибка создания файла: %v", err))
		return file.ID.Hex(), err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", fmt.Errorf("failed to convert ObjectID")
	}

	return oid.Hex(), nil
}

func (s *Storage) FindOneF(ctx context.Context, id string) (models.Files, error) {
	var file models.Files
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return file, fmt.Errorf("ошибка преобразования hex в ObjectID, hex: %s", id)
	}

	result := s.db.Collection(cloudDisk).FindOne(ctx, bson.M{"_id": oid})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return file, result.Err()
		}
		return file, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&file); err != nil {
		return file, fmt.Errorf("failed to decode user (id: %s) from DB due to error: %v", id, err)
	}

	return file, nil
}

func (s *Storage) FindFiles(ctx context.Context, file *models.Files) ([]models.Files, error) {
	// TODO неработает
	var files []models.Files
	result, err := s.db.Collection(cloudDisk).Find(ctx, bson.M{
		"userID":    file.UserID,
		"parrentID": file.ParrentID,
	})
	if err != nil {
		log.Println(err)
		return files, err
	}
	if err = result.All(ctx, &files); err != nil {
		log.Println(err)
		return files, err
	}
	return files, err
}
