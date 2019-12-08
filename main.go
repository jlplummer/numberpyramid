package main

import (
	"bufio"
	"fmt"
	"jlplummer/numberpyramid/pyramid"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/pyramid-template.html")
	t.Execute(w, currentPyramid)
}

var currentPyramid = pyramid.Pyramid{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number of rows: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//TODO: make this OS agnostic
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	userInt64, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		fmt.Println("Cannot determine number based on your entry.", err)
	}

	var userInt int = int(userInt64)
	currentPyramid.GeneratePyramid(userInt)

	fmt.Println("pyramid", currentPyramid.Pyramid)
	//fmt.Println("pyramid size", currentPyramid.PyramidSize())
	fmt.Println("pyramid hidden", currentPyramid.HiddenRows)
	fmt.Println("pyramid reversed", currentPyramid.ReversePyramid)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
