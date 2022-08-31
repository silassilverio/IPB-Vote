package routes

import (
	"net/http"

	"github.com/silassilverio/IPB-Vote/src/controllers"
)

func HanleRequest() {
	http.HandleFunc("/", controllers.Home)
}
