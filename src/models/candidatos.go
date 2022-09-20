package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Candidato struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
}

func ValidaNomeDoCandidatos(candidato *Candidato) error {
	if err := validator.Validate(candidato); err == nil {
		return err
	}
	return nil
}
