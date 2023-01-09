package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"e-menu/entity"
	"e-menu/function"
	FoodService "e-menu/service/food"

	"github.com/gorilla/mux"
)

type foodController struct{}

var (
	foodService FoodService.FoodService
)

// injection for stability
// func NewFoodController(service service.PostService) PostController {
func NewFoodController(Service FoodService.FoodService) FoodController {
	foodService = Service
	return &foodController{}

}

func (*foodController) StoreFood(resp http.ResponseWriter, req *http.Request) {

	var food entity.Food
	err := json.NewDecoder(req.Body).Decode(&food)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	foodResult, err := foodService.AddNewFood(&food)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", foodResult)

}

func (*foodController) EditFood(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	food_id := params["food_id"]

	foodID, err := strconv.ParseUint(food_id, 10, 32)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	var foodRequest entity.Food

	json.NewDecoder(req.Body).Decode(&foodRequest)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "Parsing Error", nil)
		return
	}

	foodResult, err := foodService.EditFoodById(uint(foodID), foodRequest)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", foodResult)
}

func (*foodController) DestroyFood(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	food_id := params["food_id"]

	foodID, err := strconv.Atoi(food_id)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	foodResult, err := foodService.DestroyFoodById(foodID)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", foodResult)
}

func (*foodController) ShowFood(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	food_id := params["food_id"]

	foodID, err := strconv.Atoi(food_id)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	foodResult, err := foodService.ShowFoodById(foodID)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", foodResult)
}
