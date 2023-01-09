package routes

import (
	MenuController "e-menu/controller/menu"
	router "e-menu/http"

	MenuRepository "e-menu/repository/menu"
	MenuService "e-menu/service/menu"
)

var (
	menuRepository MenuRepository.MenuRepository = MenuRepository.NewmenuRepository()
	menuService    MenuService.MenuService       = MenuService.NewMenuService(menuRepository)    //apa yang diolah (business logic)
	menuController MenuController.MenuController = MenuController.NewMenuController(menuService) //apa yang akan dikirimkan ke user (scope)
)

type MenuRoutes struct{}

func (c *MenuRoutes) Routing(httpRouter router.Router) {
	httpRouter.GET("/menus", menuController.IndexMenu)
	httpRouter.GET("/menu/{menu_id}", menuController.ShowMenu)
	httpRouter.GET("/menu/{menu_id}/detail", menuController.ShowMenu)
	httpRouter.PUT("/menu/{menu_id}", menuController.EditMenu)
	httpRouter.POST("/menu", menuController.StoreMenu)
	httpRouter.DELETE("/menu/{menu_id}", menuController.DestroyMenu)

}
