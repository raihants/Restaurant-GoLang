package controller

import "net/http"

type MenuController interface {
	IndexMenu(resp http.ResponseWriter, req *http.Request)
	ShowMenu(resp http.ResponseWriter, req *http.Request)
	StoreMenu(resp http.ResponseWriter, req *http.Request)
	EditMenu(resp http.ResponseWriter, req *http.Request)
	DestroyMenu(resp http.ResponseWriter, req *http.Request)
}
