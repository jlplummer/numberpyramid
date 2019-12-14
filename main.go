//numberpyramid presents a pyramid math problem via http to a browser
package main

import (
	"jlplummer/numberpyramid/pyramid"
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var userInt int = 5
	var currentPyramid = pyramid.Pyramid{}

	currentPyramid.GeneratePyramid(userInt)

	t, _ := template.ParseFiles("web/pyramid-template.html")
	t.Execute(w, currentPyramid)
}

func icoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", icoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
