package repository

import (
	"e-menu/entity"
)

type FoodRepository interface {
	CreateFood(food *entity.Food) (*entity.Food, error)
	UpdateFood(food *entity.Food) (*entity.Food, error)
	DeleteFoodByID(foodID int) (bool, error)
	GetFoodByID(foodID int) (entity.Food, error)
}
