package service

import (
	"e-menu/entity"
)

type RestaurntService interface {
	CreateRestaurant(post *entity.Restaurant) (*entity.Restaurant, error)
	FindAllRestaurant() ([]entity.Restaurant, error)
	ShowRestaurntById(restaurantID int) (*entity.Restaurant, error)
}
