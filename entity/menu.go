package entity

type Menu struct {
	Model
	RestaurantID int
	Title        string
	Foods        []Food
}
