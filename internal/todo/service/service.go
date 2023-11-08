package service

import (
	"github.com/danargh/golang-crud-apis/internal/todo/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}
