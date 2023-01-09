package service

import (
	"e-menu/entity"
	FoodRepository "e-menu/repository/food"
)

type foodService struct{}

var (
	foodRepo FoodRepository.FoodRepository
)

func NewFoodService(repository FoodRepository.FoodRepository) FoodService {
	foodRepo = repository
	return &foodService{}
}

func (*foodService) AddNewFood(food *entity.Food) (*entity.Food, error) {

	foodResult, err := foodRepo.CreateFood(food)
	if err != nil {
		return nil, err
	}

	return foodResult, nil
}

func (*foodService) EditFoodById(foodID uint, food entity.Food) (*entity.Food, error) {
	food.ID = foodID

	FoodResult, err := foodRepo.UpdateFood(&food)
	if err != nil {
		return nil, err
	}

	return FoodResult, nil

}

func (*foodService) DestroyFoodById(foodid int) (bool, error) {
	deleteFood, err := foodRepo.DeleteFoodByID(foodid)
	if err != nil {
		return false, err
	}

	return deleteFood, nil
}

func (*foodService) ShowFoodById(foodID int) (*entity.Food, error) {
	food, err := foodRepo.GetFoodByID(foodID)

	return &food, err

}
