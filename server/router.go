package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Phati/demo-load-balancer/db"
	"github.com/Phati/demo-load-balancer/domain"
	"github.com/gorilla/mux"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/users", CreateUserHandler()).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", GetUserHandler()).Methods(http.MethodGet)
	return
}

func CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		user := domain.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`"err":"invalid request"`))
			return
		}

		store := &db.PgStore{}

		newUser, err := store.InsertUser(r.Context(), &user, db.GetDB())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`"err":"something went wrong"`))
		}

		resp, err := json.Marshal(newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`"err":"something went wrong"`))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))

	}
}

func GetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`"err":"invalid request"`))
			return
		}

		store := &db.PgStore{}

		newUser, err := store.GetUser(r.Context(), id, db.GetDBByID(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`"err":"something went wrong"`))
		}

		resp, err := json.Marshal(newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`"err":"something went wrong"`))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))

	}
}
