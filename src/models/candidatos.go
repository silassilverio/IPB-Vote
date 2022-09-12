package models

import (
	"github.com/ipbproject/IPB-Vote/src/database"
)

type Candidatos struct {
	Id   int
	Nome string
}

func InsertNewCandidate(nome string) {
	db := database.ConectToDatabase()

	insertCandidate, err := db.Prepare("insert into candidatos (nome) values ($1)")
	if err != nil {
		panic(err.Error())
	}

	insertCandidate.Exec(nome)
	defer db.Close()
}
