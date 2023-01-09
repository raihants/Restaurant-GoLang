package repository

import (
	"e-menu/entity"
)

type MenuRepository interface {
	GetAllMenu() ([]entity.Menu, error)
	AddMenu(menu *entity.Menu) (*entity.Menu, error)
	GetMenuByID(menuID int) (entity.Menu, error)
	DeleteMenuByID(menuID int) (bool, error)
	UpdateMenu(menu *entity.Menu) (*entity.Menu, error)
	GetFoodbyMenuID(menuID int) (*entity.Menu, error)
}
