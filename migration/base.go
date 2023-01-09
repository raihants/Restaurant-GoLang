package migration

import (
	"net/http"

	"e-menu/datasource"
	"e-menu/entity"
	"e-menu/function"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type restaurant entity.Restaurant
type menu entity.Menu
type food entity.Food

func Migrate(w http.ResponseWriter, r *http.Request) {
	autoMigrate()

	function.SendResponse(w, http.StatusOK, "Success", nil)
}

// autoMigrate ...
func autoMigrate() {
	db := datasource.OpenDB()
	db.AutoMigrate(
		&restaurant{}, &menu{}, &food{},
	)
}
