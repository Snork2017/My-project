package main

import (
	"net/http"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

type record struct{
	id int
	name string
	phone int
}
func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println(name)
	age := r.URL.Query().Get("age")
	fmt.Println(age)
	w.Write([]byte("228"))
	show(*db, "privet")
}
const DB_CONNECT_STRING =
	"host=localhost port=5432 user=postgres password=2003 dbname=postgres sslmode=disable"
var db *sql.DB
func main() {
	db, err := sql.Open("postgres", DB_CONNECT_STRING)
	defer db.Close()

	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		panic("Database error")
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
		`phonebook("id" SERIAL PRIMARY KEY,` +
		`"name" varchar(50), "phone" varchar(100))`)
	fmt.Println(err)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":30097", nil)

}
func show(db sql.DB, arg string) ([]record, error) {
	var s string
	if len(arg) != 0 {
		s = "WHERE name LIKE '%" + arg + "%'"
	}
	rows, err := db.Query("SELECT * FROM phonebook " + s + " ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rs = make([]record, 0)
	var rec record
	for rows.Next() {
		err = rows.Scan(&rec.id, &rec.name, &rec.phone)
		if err != nil {
			return nil, err
		}
		rs = append(rs, rec)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	fmt.Println(rec)
	return rs, nil
}