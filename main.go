package main

import (
	"bufio"
	"fmt"
	"jlplummer/numberpyramid/pyramid"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//var pyramid = make([][]int, 5)

func generateHiddenCells(pyramid *[][]int, pyramidLength int) [][]int {
	var hiddenPyramid = make([][]int, pyramidLength)
	var hiddenPerRow = make([]int, pyramidLength)

	// instead of randomly hiding for each row...
	// randomly decide if row one has any visible cells and then go from there?
	// randomly decide if current cell's neighbor (up or left or right or down) is populated
	//       break/continue once you do something on a given row
	for g := 0; g < pyramidLength; g++ {
		rand.Seed(time.Now().UnixNano())
		hiddenPerRow[g] = rand.Intn(len((*pyramid)[g]))
	}

	fmt.Println("hiddenPerRow[]", hiddenPerRow)

	// if the cell is -1, it's "hidden"
	for j, g := range *pyramid {
		// g = slice of pyramid[j] at this point
		hiddenPyramid[j] = make([]int, len(g))

		for a, _ := range g {
			//fmt.Println("j,g,a", j, g, a)
			hiddenPyramid[j][a] = -1
		}
	}
	return hiddenPyramid
}

/*
func handler(w http.ResponseWriter, r *http.Request) {
	pyramidString := []string{}
	for j := range pyramid {
		for g := range pyramid[j] {
			pyramidString = append(pyramidString, strconv.Itoa(pyramid[j][g]))
		}
	}
	fmt.Fprintf(w, strings.Join(pyramidString, " "), r.URL.Path)
	//fmt.Println("pyramidString", pyramidString)
	//fmt.Println("pyramid", pyramid)
}
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number of rows: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	//fmt.Println(text)

	userInt64, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		fmt.Println("Cannot determine number based on your entry.", err)
	}

	var userInt int = int(userInt64)
	currentPyramid := pyramid.Pyramid{}
	currentPyramid.GeneratePyramid(userInt)

	fmt.Println("pyramid", currentPyramid)
	fmt.Println("pyramid size", currentPyramid.PyramidSize())
	//fmt.Println("pyramid hidden", generateHiddenCells(&pyramid, userInt))

	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
