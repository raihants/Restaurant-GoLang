package service

import (
	"e-menu/entity"
	repository "e-menu/repository/restaurant"
)

type service struct{}

var (
	restaurantRepository repository.RestaurantRepository
)

func NewRestaurntService(repository repository.RestaurantRepository) RestaurntService {
	restaurantRepository = repository
	return &service{}
}

func (*service) CreateRestaurant(post *entity.Restaurant) (*entity.Restaurant, error) {

	postResult, err := restaurantRepository.SaveRestaurant(post)
	if err != nil {
		return nil, err
	}

	return postResult, nil
}

func (*service) FindAllRestaurant() ([]entity.Restaurant, error) {
	restaurants, err := restaurantRepository.FindAllRestaurant()
	return restaurants, err
}

func (*service) ShowRestaurntById(restaurantID int) (*entity.Restaurant, error) {
	restaurant, err := restaurantRepository.FindRestaurantbyID(restaurantID)
	return restaurant, err

}
