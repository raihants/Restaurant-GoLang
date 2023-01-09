package entity

type Restaurant struct {
	Model
	RestaurantName string
	Menus          []Menu
}
