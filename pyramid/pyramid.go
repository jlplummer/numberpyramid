// Package generates a pyramid number structure, including data for only
// showing necessary numbers to solve the math problem.
//
// Pyramid Structure stored in top-down map
//            [0]
//          [1]  [2]
//        [3]  [4]  [5]
//     [6]  [7]  [8]  [9]
//  [10] [11] [12] [13] [14]
package pyramid

import (
	"math/rand"
	"time"
)

type Pyramid struct {
	Pyramid            [][]int
	HiddenRows         [][]int
	ReversePyramidData map[int]map[string]int
}

func (p *Pyramid) GeneratePyramid(userInt int) *Pyramid {
	p.populateCellsReverse(userInt)
	p.determineHidden(userInt)
	p.reversePyramidData(userInt)
	return p
}

func (p *Pyramid) reversePyramidData(userInt int) {
	rpd := make(map[int]map[string]int, userInt)
	cellId := 0
	for j := 0; j < userInt; j++ {
		for g := 0; g < len(p.Pyramid[j]); g++ {
			rpd[cellId] = map[string]int{"cellId": cellId, "cellValue": p.Pyramid[j][g], "cellHidden": p.HiddenRows[j][g], "cellBreak": 0}
			cellId++
		}
		rpd[cellId] = map[string]int{"cellId": cellId, "cellValue": 0, "cellHidden": 0, "cellBreak": 1}
		cellId++
	}
	p.ReversePyramidData = rpd
}

func (p *Pyramid) PyramidSize() int {
	var cells int = 0
	for x := 0; x < len((p.Pyramid)); x++ {
		cells += len((p.Pyramid)[x])
	}
	return cells
}

func pyramidCell(a, b int) int {
	return a + b
}

func (p *Pyramid) populateCellsReverse(userInt int) {
	pyramid := make([][]int, userInt)
	cntr := 0
	for x := (userInt - 1); x >= 0; x-- {
		innerLen := userInt
		pyramid[x] = make([]int, (innerLen - cntr))
		cntr++

		for j := 0; j < len(pyramid[x]); j++ {
			if x == (userInt - 1) {
				rand.Seed(time.Now().UnixNano())
				pyramid[x][j] = rand.Intn(100)
				time.Sleep(100 * time.Millisecond)
			} else {
				firstNum := pyramid[x+1][j]
				secondNum := pyramid[x+1][j+1]
				pyramid[x][j] = pyramidCell(firstNum, secondNum)
			}
		}
	}
	p.Pyramid = pyramid
}

func (p *Pyramid) determineHidden(userInt int) {
	//TODO: this actually isn't determining hidden... it's determining what to show. fix names
	pattern5 := [][]int{{10, 11, 8, 3, 5}, {13, 14, 7, 3, 5}}

	rand.Seed(time.Now().UnixNano())
	patternToUse := rand.Intn(2)

	hiddenRows := make([][]int, userInt)
	cntr := 0

	// if the cell is 1, it's "shown"
	for j, g := range p.Pyramid {
		hiddenRows[j] = make([]int, len(g))
		for a, _ := range g {
			for _, l := range pattern5[patternToUse] {
				if l == cntr {
					hiddenRows[j][a] = 1
				}
			}
			cntr++
		}
	}
	p.HiddenRows = hiddenRows
}
