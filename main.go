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

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
