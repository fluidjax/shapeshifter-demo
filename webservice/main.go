package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type status struct {
	Timestamp time.Time `json:"timeStamp"`
}

func getStatus(w http.ResponseWriter, r *http.Request) {

	var s status
	s.Timestamp = time.Now()
	log.Println("Request fulfilled at ", s.Timestamp)
	json.NewEncoder(w).Encode(s)

}

func main() {

	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	router := mux.NewRouter()
	router.HandleFunc("/status", getStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
