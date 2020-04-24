package api

import (
	"encoding/json"
	"net/http"

	"example.com/hercules/model"
	"github.com/gorilla/mux"
)

// Routes sdss
func Routes() {
	router := mux.NewRouter()

	router.HandleFunc("/", base).Methods("GET")
	router.HandleFunc("/customer", getPosts).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func base(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("BASE")
}

// AddCustomer adds a customer
func AddCustomer(w http.ResponseWriter, r *http.Request) {

	customer := model.Customer{}
	json.Unmarshal(r.Body.Read(), &customer)

}
