package service

import "github.com/DeS313/cloud-disk/internal/storage"

type Service struct {
	storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
