package db

import (
	"database/sql"
	"fmt"
	"log"
)

func StartDB(user, pass, db, host string) *sql.DB {
	DB_CONNECT_STRING := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, pass, db)
	dbConn, err := sql.Open("postgres", DB_CONNECT_STRING)
	if err != nil {
		log.Fatalf("Database opening error -->%v\n", err)
	}
	return dbConn
}
