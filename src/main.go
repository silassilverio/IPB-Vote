package main

import (
	"github.com/ipbproject/IPB-Vote/src/database"
	"github.com/ipbproject/IPB-Vote/src/routes"
)

func main() {
	database.ConectaBancoDeDados()
	routes.HanleRequest()
}
