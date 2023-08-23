package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgraph-io/badger/v4"
	"github.com/gorilla/mux"
	"github.com/meeron/honey-badger/db"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from honey badger"))
}

func handleGetDbStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	dbs, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(dbs.Stats())
}

func handleGetValue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()

	dbs, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = dbs.WriteValue(query.Get("key"), w)
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
}

func handleSetValue(w http.ResponseWriter, r *http.Request) {
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
}

func handleGetDbs(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	encoder.Encode(db.GetAll())
}

func handleDropDb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.Drop(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}

func handleDbSync(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	dbs, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = dbs.Sync()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}