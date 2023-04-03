package service

import (
	"context"

	"github.com/DeS313/cloud-disk/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) FindOne(ctx context.Context, key, value string) (models.User, error) {

	return s.storage.FindOne(ctx, value)
}

func (s *Service) Create(ctx context.Context, user *models.User) (primitive.ObjectID, error) {
	//  TODO придумать как проверять если ли пользователь с таким email
	return primitive.NewObjectID(), nil
}
