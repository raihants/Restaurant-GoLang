package service

import (
	"e-menu/entity"
)

type FoodService interface {
	// Validate(post *entity.Post) error
	AddNewFood(food *entity.Food) (*entity.Food, error)
	EditFoodById(foodID uint, food entity.Food) (*entity.Food, error)
	DestroyFoodById(foodID int) (bool, error)
	ShowFoodById(foodID int) (*entity.Food, error)
}
