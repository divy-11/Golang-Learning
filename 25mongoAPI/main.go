package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/divy-11/mongoapi/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":1717", r))
	fmt.Println("Try hitting Port 1717 !!")
}
