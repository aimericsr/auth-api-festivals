package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aimericsr/auth-api-festivals/util"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}

	_, err = sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	//router.HandleFunc("/", middleware.EnsureValidToken()(Root)).Methods("GET")
	router.HandleFunc("/v1/health", Root).Methods("GET")
	http.ListenAndServe(config.ServerAddress, router)
}

func Root(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Root API Endpoint")
}
