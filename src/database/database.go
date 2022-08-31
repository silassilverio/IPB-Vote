package database

import (
	"database/sql"
	"fmt"
	"os"
)

func ConectToDatabase() *sql.DB {
	conectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	conection := conectionString
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
