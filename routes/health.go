package routes

import (
	router "e-menu/http"
	"fmt"
	"net/http"
)

var (
	httpRouterHealth router.Router = router.NewMuxRouter()
)

type HealthRoutes struct{}

func (c *HealthRoutes) Routing() {
	httpRouterHealth.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up n run")
	})
}
