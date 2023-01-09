package repository

import (
	"e-menu/datasource"
	"e-menu/entity"
)

type dbrepo struct{}

func NewDatabaseRepository() RestaurantRepository {
	return &dbrepo{}
}

func (*dbrepo) FindAllRestaurant() ([]entity.Restaurant, error) {
	db := datasource.OpenDB()

	var restaurant []entity.Restaurant

	db.Find(&restaurant)

	defer datasource.CloseDB()

	return restaurant, nil

}

func (*dbrepo) SaveRestaurant(restaurant *entity.Restaurant) (*entity.Restaurant, error) {
	db := datasource.OpenDB()

	db.Create(&restaurant)

	defer datasource.CloseDB()
	return restaurant, nil
}

func (*dbrepo) FindRestaurantbyID(restaurantID int) (*entity.Restaurant, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	var restaurant entity.Restaurant
	db.First(&restaurant, restaurantID)

	return &restaurant, nil

}
