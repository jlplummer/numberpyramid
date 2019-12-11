package main

import (
	"fmt"
	"jlplummer/numberpyramid/pyramid"
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var userInt int = 5
	var currentPyramid = pyramid.Pyramid{}

	currentPyramid.GeneratePyramid(userInt)

	fmt.Println("pyramid", currentPyramid.Pyramid)
	//fmt.Println("pyramid size", currentPyramid.PyramidSize())
	fmt.Println("pyramid hidden", currentPyramid.HiddenRows)
	fmt.Println("pyramid reversed", currentPyramid.ReversePyramid)

	t, _ := template.ParseFiles("web/pyramid-template.html")
	t.Execute(w, currentPyramid)
}

func icoHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", icoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
