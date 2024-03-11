package service

import (
	"github.com/daniil-sargsyan/go-basic/pkg/storage"
)

type Service interface {
	//add methods to service and implement like this
}

type service struct {
	db storage.Storage
}

func NewService(store storage.Storage) Service {
	return &service{db: store}
}
