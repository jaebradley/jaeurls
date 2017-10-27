package jaeurls

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/jaeurls", createUrl).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func createUrl(w http.ResponseWriter, r *http.Request) {
	log.Print(w)
	log.Print(r)
}
