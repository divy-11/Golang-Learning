package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleReq).Methods("GET")

	fmt.Println("Listening on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang.</h1>")) //as everything in web is dealt in form of bytes
}
