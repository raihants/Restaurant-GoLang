package repository

import (
	"e-menu/entity"
)

type RestaurantRepository interface {
	FindAllRestaurant() ([]entity.Restaurant, error)
	SaveRestaurant(post *entity.Restaurant) (*entity.Restaurant, error)
	FindRestaurantbyID(restaurantID int) (*entity.Restaurant, error)
}
