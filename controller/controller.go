package controller

import (
	"SI/repository"
	"SI/service"

	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
	R  *repository.Repository
	S  *service.Service
}

func NewController(db *gorm.DB, r *repository.Repository, s *service.Service) *Controller {
	return &Controller{
		DB: db,
		R:  r,
		S:  s,
	}
}
