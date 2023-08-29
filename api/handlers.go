package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meeron/honey-badger/db"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from honey badger"))
}

func handleGetDbStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(db.Stats())
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

	db, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = db.Sync()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}

func handleDeleteWithKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()

	db, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = db.DeleteByKey(query.Get("key"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}

func handleDeleteWithPrefix(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query := r.URL.Query()
	prefix := query.Get("prefix")

	if prefix == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Prefix cannot be empty"))
		return
	}

	db, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = db.DeleteByPrefix(prefix)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}

func handleCreateDb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	options := db.NewDbOptions{}

	if err := decoder.Decode(&options); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := options.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	db, err := db.Create(options)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encoder := json.NewEncoder(w)

	err = encoder.Encode(db.Stats())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func handleGetByPrefix(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query := r.URL.Query()
	prefix := query.Get("prefix")

	if prefix == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Prefix cannot be empty"))
		return
	}

	db, err := db.Get(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//w.Header().Set("Content-Type", "text/plain")
	if err := db.GetByPrefix(prefix, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
