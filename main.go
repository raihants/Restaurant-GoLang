package main

import (
	router "e-menu/http"
	"e-menu/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var (
	// dbRepository   repository.PostRepository = repository.NewDatabaseRepository()
	// postService    service.PostService       = service.NewPostService(dbRepository)
	// postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewMuxRouter()
)

// @title Garuda Express Delivery
// @version 1.0.0
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	httpRouter.BASE("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Healthy")
	})

	// for migration purpose
	migrateRoutes := routes.MigrateRoutes{}
	migrateRoutes.Routing(httpRouter)

	foodRoutes := routes.FoodRoutes{}
	foodRoutes.Routing(httpRouter)

	menuRoutes := routes.MenuRoutes{}
	menuRoutes.Routing(httpRouter)

	restaurantRoutes := routes.RestaurantRoutes{}
	restaurantRoutes.Routing(httpRouter)

	httpRouter.SERVE()

}
