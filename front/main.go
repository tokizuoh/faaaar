package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/playground", func(rw http.ResponseWriter, r *http.Request) {
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
