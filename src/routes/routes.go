package routes

import (
	"net/http"

	"github.com/ipbproject/IPB-Vote/src/controllers"
)

func HanleRequest() {
	http.HandleFunc("/", controllers.Home)
}
