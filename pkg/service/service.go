package service

import "github.com/alexkozopolianski/todo-app/pkg/repository"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
