package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("*.html"))
	u := usuario{"Maria", "maria@hotmail.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}

type usuario struct {
	Nome  string
	Email string
}

func main() {

	http.HandleFunc("/home", home)

	fmt.Println("Listening at port: 5000. Waiting requests...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
