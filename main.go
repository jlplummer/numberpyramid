package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var pyramid = make([][]int, 5)

func pyramidCell(a, b int) int {
	return a + b
}

func generatePyramid(userInt int) [][]int {
	/*
		[0] = [3, 5, 10, 10]
		[1] = [8, 15, 20]
		[2] = [23, 35]
		[3] = [58]
	*/
	//pyramid := make([][]int, userInt)
	for x := 0; x < userInt; x++ {
		innerLen := userInt
		pyramid[x] = make([]int, (innerLen - x))
		for j := 0; j < len(pyramid[x]); j++ {
			if x == 0 {
				rand.Seed(time.Now().UnixNano())
				pyramid[x][j] = rand.Intn(100)
				time.Sleep(100 * time.Millisecond)
			} else {
				firstNum := pyramid[x-1][j]
				secondNum := pyramid[x-1][j+1]
				pyramid[x][j] = pyramidCell(firstNum, secondNum)
			}
		}
	}

	return pyramid
}

func pyramidSize(pyramid *[][]int) int {
	var cells int = 0
	for x := 0; x < len((*pyramid)); x++ {
		// de-reference the pyramid, then get the slice
		cells += len((*pyramid)[x])
	}
	return cells
}

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
	pyramid := generatePyramid(userInt)

	fmt.Println("pyramid", pyramid)
	//fmt.Println("pyramid size", pyramidSize(&pyramid))
	fmt.Println("pyramid hidden", generateHiddenCells(&pyramid, userInt))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
