package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Run(addr string) error {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHome).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}", handleGetDbStats).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}/get", handleGetValue).Methods(http.MethodGet)
	router.HandleFunc("/dbs/{name}/set", handleSetValue).Methods(http.MethodPost)

	return http.ListenAndServe(addr, router)
}
