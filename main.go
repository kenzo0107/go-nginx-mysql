package main

import (
	"app/config"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// NewDB return database global connection handle.
func NewDB(conf config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Handler ... DB Handler
type Handler struct {
	DB *sql.DB
}

func main() {
	var err error

	appMode := os.Getenv("GO_ENV")
	if appMode == "" {
		panic("failed to get application mode, check whether GO_ENV is set.")
	}

	conf, err := config.NewConfig(appMode)
	if err != nil {
		panic(err.Error())
	}

	db, err := NewDB(conf.DB)
	if err != nil {
		panic(err.Error())
	}

	h := Handler{DB: db}
	r := mux.NewRouter()

	r.Methods("GET").Path("/").HandlerFunc(h.top)
	r.Methods("GET").Path("/article").HandlerFunc(h.article)

	fmt.Fprint(os.Stdout, ">> Start to listen http server post :8080\n")
	if err = http.ListenAndServe(":8080", r); err != nil {
		panic(err.Error())
	}
}

func (h *Handler) top(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("select id, name from users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err.Error())
		}
		fmt.Fprintf(w, "id: %d name: %s\n", id, name)
	}
}

func (h *Handler) article(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Article !\n")
}
