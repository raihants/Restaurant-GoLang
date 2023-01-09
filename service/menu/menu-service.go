package service

import (
	"e-menu/entity"
	MenuRepository "e-menu/repository/menu"
)

type menuService struct{}

var (
	MenuRepo MenuRepository.MenuRepository
)

func NewMenuService(repository MenuRepository.MenuRepository) MenuService {
	MenuRepo = repository
	return &menuService{}
}

func (*menuService) ShowAllMenu() ([]entity.Menu, error) {
	lobbies, err := MenuRepo.GetAllMenu()
	return lobbies, err
}
func (*menuService) ShowMenuById(menuID int) (*entity.Menu, error) {

	menu, err := MenuRepo.GetMenuByID(menuID)

	return &menu, err
}
func (*menuService) AddMenu(menu *entity.Menu) (*entity.Menu, error) {
	MenuResult, err := MenuRepo.AddMenu(menu)
	if err != nil {
		return nil, err
	}

	return MenuResult, nil
}
func (*menuService) EditMenuById(menuID uint, menu entity.Menu) (*entity.Menu, error) {

	menu.ID = menuID

	MenuResult, err := MenuRepo.UpdateMenu(&menu)
	if err != nil {
		return nil, err
	}

	return MenuResult, nil
}
func (*menuService) DestroyMenuById(menuID int) (bool, error) {
	deleteMenu, err := MenuRepo.DeleteMenuByID(menuID)
	if err != nil {
		return false, err
	}

	return deleteMenu, nil
}

func (*menuService) ShowMenuDetail(menuID int) (*entity.Menu, error) {

	MenuDetailResult, err := MenuRepo.GetFoodbyMenuID(menuID)
	if err != nil {
		return nil, err
	}

	return MenuDetailResult, nil
}
