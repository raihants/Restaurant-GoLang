package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"e-menu/entity"
	"e-menu/function"

	MenuService "e-menu/service/menu"

	"github.com/gorilla/mux"
)

type menuController struct{}

var (
	menuService MenuService.MenuService
)

// injection for stability
// func NewUserController(service service.PostService) PostController {
func NewMenuController(service MenuService.MenuService) MenuController {
	menuService = service
	return &menuController{}

}

func (*menuController) IndexMenu(resp http.ResponseWriter, req *http.Request) {
	lobbies, err := menuService.ShowAllMenu()
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}
	_ = function.SendResponse(resp, http.StatusOK, "success", lobbies)

}

func (*menuController) StoreMenu(resp http.ResponseWriter, req *http.Request) {

	var menu entity.Menu
	err := json.NewDecoder(req.Body).Decode(&menu)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	menuResult, err := menuService.AddMenu(&menu)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", menuResult)

}

func (*menuController) ShowMenu(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	menu_id := params["menu_id"]

	menuID, err := strconv.Atoi(menu_id)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	menuResult, err := menuService.ShowMenuById(menuID)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", menuResult)

}

func (*menuController) EditMenu(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	menu_id := params["menu_id"]

	menuID, err := strconv.ParseUint(menu_id, 10, 32)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	var menuRequest entity.Menu

	json.NewDecoder(req.Body).Decode(&menuRequest)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "Parsing Error", nil)
		return
	}

	menuResult, err := menuService.EditMenuById(uint(menuID), menuRequest)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", menuResult)

}

func (*menuController) DestroyMenu(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	menu_id := params["menu_id"]

	menuID, err := strconv.Atoi(menu_id)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	menuResult, err := menuService.DestroyMenuById(menuID)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", menuResult)

}
