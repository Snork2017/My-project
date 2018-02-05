package api

import (
	"github.com/go-chi/chi"
	"net/http"
	"log"
	"encoding/json"
	"My-project/models"
)

func (api *API) initUsersRoutes(r chi.Router) {
	r.Get("/", api.getUsersByName)
}

func (api *API) getUsersByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := api.showUsersByName(name)
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

func (api *API) showUsersByName(name string) ([]models.User, error) {
	s := "WHERE name LIKE '%" + name + "%'"
	rows, err := api.db.Query("SELECT * FROM phonebook " + s + " ORDER BY id")
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
		if err := rows.Scan(&user.ID, &user.Name, &user.Phone); err != nil {
			log.Println("rows.Scan():", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}