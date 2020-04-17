package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StarrtAPI starts hercules API
func StartAPI() {
	router := setupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/hercules").
		HandlerFunc(getHercules)
	return router
}

func getHercules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("HERCULES!")
	w.WriteHeader(http.StatusOK)
}
