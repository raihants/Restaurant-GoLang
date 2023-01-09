package repository

import (
	"e-menu/datasource"
	"e-menu/entity"
)

type foodRepository struct{}

func NewFoodRepository() FoodRepository {

	return &foodRepository{}

}

func (*foodRepository) CreateFood(food *entity.Food) (*entity.Food, error) {

	db := datasource.OpenDB()

	db.Create(&food)

	defer datasource.CloseDB()
	return food, nil
}

func (*foodRepository) UpdateFood(food *entity.Food) (*entity.Food, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	err := db.Save(&food).Error
	if err != nil {
		return nil, err
	}

	return food, nil

}

func (*foodRepository) DeleteFoodByID(foodID int) (bool, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	err := db.Unscoped().Where("id = ?", foodID).Delete(entity.Food{}).Error
	if err != nil {
		return false, err
	}

	return true, nil

}

func (*foodRepository) GetFoodByID(foodID int) (entity.Food, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	var food entity.Food
	db.First(&food, foodID)

	return food, nil

}
