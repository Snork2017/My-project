package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"log"
	"encoding/json"
	"My-project/models"
)


func (api *API) initUsersRoutes(r chi.Router) {
	r.Get("/", api.getUsersByName)
	r.Get("/createusers", api.insert)
}

func (api *API) getUsersByName(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	if firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	users, err := api.showUsersByName(firstname)
	if err != nil {
		log.Println("show():", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersResponse, err := json.Marshal(users)
	if err != nil {
		log.Println("json.Marshal(users):", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(usersResponse); err != nil {
		log.Println("w.Write():", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//noinspection GoUnresolvedReference
func (api *API) showUsersByName(name string) ([]models.User, error) {
	s := "WHERE firstname LIKE '%" + name + "%'"
	rows, err := api.db.Query("SELECT * FROM users " + s + " ORDER BY id")
	if err != nil {
		log.Println("db.Query():", err)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("rows.Close():", err)
		}
	}()

	var user models.User
	var users []models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Secondname, &user.Thirdname,  &user.Phone); err != nil {
			log.Println("rows.Scan():", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (api *API) insert(w http.ResponseWriter, r *http.Request){

	firstname := r.URL.Query().Get("firstname")
	if firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	secondname := r.URL.Query().Get("secondname")
	if secondname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	thirdname := r.URL.Query().Get("thirdname")
	if thirdname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err :=  api.db.Exec("INSERT INTO users VALUES (default, $1, $2, $3, $4)",
		firstname, secondname, thirdname, phone)
	if err != nil {
		return
	}

	w.Write([]byte(fmt.Sprintf("Уважаемый %s приветствуем вас", firstname)))
}


