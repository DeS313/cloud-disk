package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/DeS313/cloud-disk/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) FindOneFile(ctx context.Context, id string) (models.Files, error) {
	return s.storage.FindOneF(ctx, id)
}

func (s *Service) FindFile(ctx context.Context, userID, parentID string) ([]models.Files, error) {

	return s.storage.FindFiles(ctx, userID, parentID)
}

func (s *Service) CreateFile(ctx context.Context, file *models.Files) (string, error) {
	fmt.Println(file.ParrentID, "parrentID")
	parent, err := s.storage.FindOneF(ctx, file.ParrentID.Hex())
	file.ID = primitive.NewObjectID()

	if errors.Is(err, mongo.ErrNoDocuments) {
		file.ParrentID = file.UserID
		file.Path = file.Name
		log.Println(err, "mongo")
	} else {
		file.ParrentID = parent.ID
		file.Path = fmt.Sprintf("%s/%s", parent.Path, file.Name)
		fmt.Println(file.ParrentID, "parrentID")
	}

	id, err := s.storage.CreateF(ctx, file)
	if err != nil {
		return "", err
	}

	if err := s.createDir(file); err != nil {
		return "", err
	}

	return id, err
}

func (s *Service) createDir(file *models.Files) error {
	filePath := fmt.Sprintf("%s/%s/%s", s.config.FilePath, file.UserID.Hex(), file.Path)
	_, err := os.Open(filePath)
	fmt.Println(err)
	if err != nil {
		if err = os.MkdirAll(filePath, 0777); err != nil {
			return fmt.Errorf("ошибка создания файла %s, path: %s", err, filePath)
		}
		return nil
	}
	return err
}

func (s *Service) CreateDir(file *models.Files) error {
	return s.createDir(file)
}

func (s *Service) CreateF(ctx context.Context, file *models.Files) (string, error) {
	file.ID = primitive.NewObjectID()
	return s.storage.CreateF(ctx, file)
}
