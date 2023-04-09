package service

import (
	"github.com/DeS313/cloud-disk/internal/config"
	"github.com/DeS313/cloud-disk/internal/storage"
)

type Service struct {
	storage *storage.Storage
	config  *config.Config
}

func NewService(storage *storage.Storage, config *config.Config) *Service {
	return &Service{
		storage: storage,
		config:  config,
	}
}
