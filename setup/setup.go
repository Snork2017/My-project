package setup

import (
	"database/sql"
	"log"
)

func CreateTable(dbCon *sql.DB) {
	_, err := dbCon.Exec("CREATE TABLE IF NOT EXISTS " +
		`users("id" SERIAL PRIMARY KEY,` +
		`"firstname" varchar(50), "secondname" varchar(50), "thirdname" varchar(50), "phone" varchar(20))`)
	if err != nil {
		log.Println("dbCon.Exec():", err)
	}
}

