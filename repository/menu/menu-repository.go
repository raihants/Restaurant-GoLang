package repository

import (
	"e-menu/datasource"
	"e-menu/entity"
)

type menuRepository struct{}

func NewmenuRepository() MenuRepository {

	return &menuRepository{}

}

func (*menuRepository) GetAllMenu() ([]entity.Menu, error) {

	db := datasource.OpenDB()
	defer db.Close()

	var lobbies []entity.Menu

	db.Find(&lobbies)

	return lobbies, nil
}

func (*menuRepository) AddMenu(Menu *entity.Menu) (*entity.Menu, error) {

	db := datasource.OpenDB()
	defer datasource.CloseDB()

	db.Create(&Menu)

	return Menu, nil
}

func (*menuRepository) GetMenuByID(menuID int) (entity.Menu, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	var menu entity.Menu
	db.First(&menu, menuID)

	return menu, nil

}

func (*menuRepository) DeleteMenuByID(menuID int) (bool, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	err := db.Unscoped().Where("id = ?", menuID).Delete(entity.Menu{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (*menuRepository) UpdateMenu(menu *entity.Menu) (*entity.Menu, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	err := db.Save(&menu).Error
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (*menuRepository) GetFoodbyMenuID(menuID int) (*entity.Menu, error) {
	db := datasource.OpenDB()
	defer datasource.CloseDB()

	var menuDetail entity.Menu

	err := db.Debug().Preload("Users").Where("id = ?", menuID).Find(&menuDetail).Error
	if err != nil {
		return nil, err
	}

	return &menuDetail, nil
}
