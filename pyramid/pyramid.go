package pyramid

import (
	"fmt"
	"math/rand"
	"time"
)

type Pyramid struct {
	pyramid    [][]int
	hiddenRows [][]int
}

func (p *Pyramid) GeneratePyramid(userInt int) *Pyramid {
	/*
		[0] = [3, 5, 10, 10]
		[1] = [8, 15, 20]
		[2] = [23, 35]
		[3] = [58]
	*/
	p.pyramid = populateCells(userInt)
	p.hiddenRows = determineHidden(p.pyramid, p.PyramidSize())
	return p
}

func (p *Pyramid) PyramidSize() int {
	var cells int = 0
	for x := 0; x < len((p.pyramid)); x++ {
		// de-reference the pyramid, then get the slice
		cells += len((p.pyramid)[x])
	}
	return cells
}

func pyramidCell(a, b int) int {
	return a + b
}

func populateCells(userInt int) [][]int {
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

//TODO: figure out syntax to pass pyramid in as a reference instead of a copy
func determineHidden(pyramid [][]int, pyramidLength int) [][]int {
	var hiddenPyramid = make([][]int, pyramidLength)
	var hiddenPerRow = make([]int, pyramidLength)

	// instead of randomly hiding for each row...
	// randomly decide if row one has any visible cells and then go from there?
	// randomly decide if current cell's neighbor (up or left or right or down) is populated
	//       break/continue once you do something on a given row

	fmt.Println("hiddenPerRow[]", hiddenPerRow)

	// if the cell is -1, it's "hidden"
	for j, g := range pyramid {
		// g = slice of pyramid[j] at this point
		hiddenPyramid[j] = make([]int, len(g))

		for a, _ := range g {
			//fmt.Println("j,g,a", j, g, a)
			hiddenPyramid[j][a] = -1
		}
	}
	return hiddenPyramid
}
