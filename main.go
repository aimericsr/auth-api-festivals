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
		//log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/test", HealthCheck).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
