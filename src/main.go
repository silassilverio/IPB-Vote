package main

import (
	"net/http"

	"github.com/silassilverio/IPB-Vote/src/routes"
)

func main() {
	routes.HanleRequest()
	http.ListenAndServe(":8000", nil)
}
