package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/student", StudentRouter).Methods("GET")
	r.HandleFunc("/api/v1/student/{name}", StudentNameRouter).Methods("GET")

	http.ListenAndServe(":3000", r)

}

func StudentRouter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Println(vars)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Mux"))

	// fmt.Fprintf(w)

}

func StudentNameRouter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Println(vars)

	w.WriteHeader(http.StatusFound)
	w.Write([]byte(vars["name"]))

}
