package main

import (
	"html/template"
	"log"
	"net/http"
)

// TODO: .envに入れる
const basicAuthUsername = "root"
const basicAuthPassword = "root"

// TODO: ユーザーごとに認証したか判別する
var doneAuth = false

func basicAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if ok == false {
		return false
	}

	if username != basicAuthUsername || password != basicAuthPassword {
		return false
	}

	doneAuth = true
	return true
}

func main() {
	// Basic Authentication
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if basicAuth(r) {
			http.Redirect(rw, r, "/playground", http.StatusMovedPermanently)
			return
		} else {
			rw.Header().Set("WWW-Authenticate", "Basic")
			rw.WriteHeader(http.StatusUnauthorized)
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})

	// GraphQL Playground
	http.HandleFunc("/playground", func(rw http.ResponseWriter, r *http.Request) {
		if doneAuth != true {
			http.Redirect(rw, r, "/", http.StatusSeeOther)
			return
		}

		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Fatal(err)
		}

		if err := t.Execute(rw, nil); err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8081", nil)
}
