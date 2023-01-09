package service

import (
	"e-menu/entity"
)

type MenuService interface {
	ShowAllMenu() ([]entity.Menu, error)
	ShowMenuById(menuID int) (*entity.Menu, error)
	AddMenu(menu *entity.Menu) (*entity.Menu, error)
	EditMenuById(menuID uint, Menu entity.Menu) (*entity.Menu, error)
	DestroyMenuById(menuID int) (bool, error)
	ShowMenuDetail(menuID int) (*entity.Menu, error)
}
