package controller

import "net/http"

type RestaurantController interface {
	GetRestaurants(resp http.ResponseWriter, req *http.Request)
	GetRestaurantbyID(resp http.ResponseWriter, req *http.Request)
	AddRestaurants(resp http.ResponseWriter, req *http.Request)
}
