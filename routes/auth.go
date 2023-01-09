package routes

import (
	router "e-menu/http"
	"fmt"
	"net/http"
)

var (
	httpRouterAuth router.Router = router.NewMuxRouter()
)

type AuthhRoutes struct{}

func (c *AuthhRoutes) Routing() {
	httpRouterAuth.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up n run")
	})
}
