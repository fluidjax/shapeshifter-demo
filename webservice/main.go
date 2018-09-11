package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	fmt.Println("server is running")
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
