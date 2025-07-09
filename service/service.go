package service

import "SI/repository"

type Service struct {
	R *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		R: r,
	}
}
