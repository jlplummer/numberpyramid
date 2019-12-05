package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	pyramid := make([][]int, userInt)
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

/*
From number pyramids website [https://nrich.maths.org/numberpyramids]
var epsilon = 0.00000001

function round(x)
{
   if(Math.abs(Math.floor(x)-x) < epsilon) return x;
   if(Math.abs(Math.floor(10*x)-10*x) < epsilon) return x.toFixed(1);
   if(Math.abs(Math.floor(100*x)-100*x)<epsilon) return x.toFixed(2);
   return (x.toFixed(2) + "...");
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
	pyramid := generatePyramid(userInt)

	fmt.Println("pyramid", pyramid)
}
