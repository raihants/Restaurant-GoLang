package controller

import "net/http"

type FoodController interface {
	StoreFood(resp http.ResponseWriter, req *http.Request)
	ShowFood(resp http.ResponseWriter, req *http.Request)
	EditFood(resp http.ResponseWriter, req *http.Request)
	DestroyFood(resp http.ResponseWriter, req *http.Request)
}
