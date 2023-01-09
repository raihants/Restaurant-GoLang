package routes

import (
	controller "e-menu/controller/restaurant"
	router "e-menu/http"
	repository "e-menu/repository/restaurant"

	service "e-menu/service/restaurant"
)

var (
	dbRepository   repository.RestaurantRepository = repository.NewDatabaseRepository()        //query ke DB
	postService    service.RestaurntService        = service.NewRestaurntService(dbRepository) //apa yang diolah (business logic)
	postController controller.RestaurantController = controller.NewPostController(postService) //apa yang akan dikirimkan ke user (scope)
)

type RestaurantRoutes struct{}

func (c *RestaurantRoutes) Routing(httpRouter router.Router) {
	httpRouter.GET("/restaurants", postController.GetRestaurants)
	httpRouter.GET("/restaurant/{restaurant_id}", postController.GetRestaurantbyID)
	httpRouter.POST("/restaurant", postController.AddRestaurants)
}
