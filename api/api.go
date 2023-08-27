package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(addr string) error {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHome).Methods(http.MethodGet)
	router.HandleFunc("/dbs", handleGetDbs).Methods(http.MethodGet)
	router.HandleFunc("/dbs", handleCreateDb).Methods(http.MethodPost)
	router.HandleFunc("/dbs/{name}", handleGetDbStats).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}", handleDropDb).Methods(http.MethodDelete)
	router.HandleFunc("/dbs/{name}/sync", handleDbSync).Methods(http.MethodPost)
	router.HandleFunc("/dbs/{name}/entry", handleGetValue).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}/entry", handleSetValue).Methods(http.MethodPost)
	router.HandleFunc("/dbs/{name}/entry", handleDeleteWithKey).Methods(http.MethodDelete)
	router.HandleFunc("/dbs/{name}/entries", handleDeleteWithPrefix).Methods(http.MethodDelete)

	log.Printf("Listening on %s...", addr)
	return http.ListenAndServe(addr, router)
}
