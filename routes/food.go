package routes

import (
	FoodController "e-menu/controller/food"
	router "e-menu/http"
	FoodRepository "e-menu/repository/food"
	FoodService "e-menu/service/food"
)

var (
	foodRepository FoodRepository.FoodRepository = FoodRepository.NewFoodRepository()            //query ke DB
	foodService    FoodService.FoodService       = FoodService.NewFoodService(foodRepository)    //apa yang diolah (business logic)
	foodController FoodController.FoodController = FoodController.NewFoodController(foodService) //apa yang akan dikirimkan ke food (scope)
)

type FoodRoutes struct{}

func (c *FoodRoutes) Routing(httpRouter router.Router) {

	httpRouter.GET("/food/{food_id}", foodController.ShowFood)
	httpRouter.POST("/food", foodController.StoreFood)
	httpRouter.PUT("/food{food_id}", foodController.EditFood)
	httpRouter.DELETE("/food", foodController.DestroyFood)
}
