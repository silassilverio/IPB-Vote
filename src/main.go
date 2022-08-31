package main

import (
	"net/http"

	"github.com/ipbproject/IPB-Vote/src/routes"
)

func main() {
	routes.HanleRequest()
	http.ListenAndServe(":8000", nil)
}
