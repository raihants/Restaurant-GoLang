package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"e-menu/entity"
	"e-menu/function"
	service "e-menu/service/restaurant"

	"github.com/gorilla/mux"
)

type restaurantController struct{}

var (
	restaurantService service.RestaurntService
)

// injection for stability
func NewPostController(service service.RestaurntService) RestaurantController {
	restaurantService = service
	return &restaurantController{}

}

func (*restaurantController) GetRestaurants(resp http.ResponseWriter, req *http.Request) {
	Posts, err := restaurantService.FindAllRestaurant()
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}
	_ = function.SendResponse(resp, http.StatusOK, "success", Posts)

}

func (*restaurantController) AddRestaurants(resp http.ResponseWriter, req *http.Request) {

	var restaurant entity.Restaurant
	err := json.NewDecoder(req.Body).Decode(&restaurant)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	restaurantResult, err := restaurantService.CreateRestaurant(&restaurant)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", restaurantResult)

}

func (*restaurantController) GetRestaurantbyID(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	restaurant_id := params["restaurant_id"]

	restaurantID, err := strconv.Atoi(restaurant_id)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	MenuResult, err := restaurantService.ShowRestaurntById(restaurantID)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", MenuResult)
}
