package main

import (
	"encoding/json"
	"net/http"

	"github.com/dgraph-io/badger/v4"
	"github.com/gorilla/mux"
	"github.com/meeron/honey-badger/db"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from honey badger"))
	}).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		dbs, err := db.Get(params["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(dbs.Stats())
	}).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}/get", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		query := r.URL.Query()

		dbs, err := db.Get(params["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		value, err := dbs.Get(query.Get("key"))
		if err == badger.ErrKeyNotFound {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(value)
	}).Methods(http.MethodGet)

	router.HandleFunc("/dbs/{name}/set", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		query := r.URL.Query()

		key := query.Get("key")
		if key == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(badger.ErrEmptyKey.Error()))
			return
		}

		dbs, err := db.Get(params["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = dbs.Set(key, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("Ok"))

	}).Methods(http.MethodPost)

	http.ListenAndServe(":8080", router)
}
