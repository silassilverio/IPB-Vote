package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ipbproject/IPB-Vote/src/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var candidato *models.Candidatos
		json.NewDecoder(r.Body).Decode(&candidato)

		models.InsertNewCandidate(candidato.Nome)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
