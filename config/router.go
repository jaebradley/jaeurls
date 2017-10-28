package config

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"net/url"
)

func StartRouter(session *mgo.Session) {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/", createUrl(session)).Methods("POST")
	r.HandleFunc("/api/v1/{jae[a-zA-Z0-9]+}", redirectUrl(session)).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

type CreateUrlData struct {
	Url string	`json: "url"`
}

type CreatedUrlData struct {
	Url string `json: "url"`
}

func createUrl(session *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var data CreateUrlData
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		u, err := url.Parse(data.Url)
		if err != nil {
			http.Error(w, "Invalid URL", 400)
		}

		key, e := InsertUrl(session, *u)

		if e != nil {
			http.Error(w, "Unable to insert URL", 500)
		}

		cu := CreateUrl(key)

		json.NewEncoder(w).Encode(CreatedUrlData{Url: cu.String()})
	}
}

func redirectUrl(session *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key, e := ParseUrl(*r.URL)

		if e != nil {
			http.Error(w, e.Error(), 400)
		}

		u, e := IdentifyUrl(session, key)

		if e != nil {
			http.Error(w, "unable to identify URL", 400)
		}

		http.Redirect(w, r, u.String(), 301)
	}
}
