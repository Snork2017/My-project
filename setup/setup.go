package setup

import (
	"database/sql"
	"log"
)

func CreateTable(dbCon *sql.DB) {
	_, err := dbCon.Exec("CREATE TABLE IF NOT EXISTS " +
		`phonebook("id" SERIAL PRIMARY KEY,` +
		`"name" varchar(50), "phone" varchar(100))`)
	if err != nil {
		log.Println("dbCon.Exec():", err)
	}
}