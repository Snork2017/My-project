package api

import (
	"My-project/conf"
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// API - type for dependency injection
type API struct {
	cnf *conf.Conf
	ctx context.Context
	db  *sql.DB

	//hc    *http.Client
	Hs *http.Server
}

// New initialize api with routes
func New(ctx context.Context, cnf conf.Conf, dbConn *sql.DB) *API {
	api := &API{
		cnf: &cnf,
		ctx: ctx,
		db:  dbConn,
	}
	api.initRouter()
	return api
}

func (api *API) initRouter() {
	r := chi.NewRouter()

	r.Use(setContentType)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", api.initUsersRoutes)

	})

	api.Hs = &http.Server{
		Addr:           api.cnf.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(api.cnf.HTTPReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(api.cnf.HTTPWriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 Mb
	}
}

func (api *API) Start() {
	log.Println("launching the My-project service at", api.cnf.Addr)
	err := api.Hs.ListenAndServe()
	if err != nil && err.Error() == "http: Server closed" {
		log.Println("api port is closed")
		return
	}
	log.Panic(err)
}

// Stop will call Shutdown function
func (api *API) Stop() error {
	ctx, cnl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cnl()
	return api.Hs.Shutdown(ctx)
}

func setContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
