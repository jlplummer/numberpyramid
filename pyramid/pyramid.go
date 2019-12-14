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
	"fmt"
	"math/rand"
	"time"
)

type Pyramid struct {
	Pyramid            [][]int
	Pyramid2           [][]int
	HiddenRows         [][]int
	ReversePyramid     [][]int
	ReversedHiddenRows [][]int
	PyramidData        map[int]map[string]int
	ReversePyramidData map[int]map[string]int
}

func (p *Pyramid) GeneratePyramid(userInt int) *Pyramid {
	p.populateCells(userInt)
	p.determineHidden(userInt)
	p.reversePyramid(userInt)
	p.reverseHidden(userInt)
	p.populatePyramidData(userInt)
	p.reversePyramidData(userInt)
	p.populateCellsReverse(userInt)
	fmt.Println("pyramid2", p.Pyramid2)
	return p
}

func (p *Pyramid) populatePyramidData(userInt int) {
	pd := make(map[int]map[string]int)
	var cellId int = 0
	for j := 0; j < userInt; j++ {
		for g := 0; g < len(p.Pyramid[j]); g++ {
			pd[cellId] = map[string]int{"cellId": cellId, "cellValue": p.Pyramid[j][g], "cellHidden": p.HiddenRows[j][g], "cellBreak": 0}
			cellId++
		}
		pd[cellId] = map[string]int{"cellId": cellId, "cellValue": 0, "cellHidden": 0, "cellBreak": 1}
		cellId++
	}

	p.PyramidData = pd
}

func (p *Pyramid) reversePyramidData(userInt int) {
	rpd := make(map[int]map[string]int)
	cellId := 0
	for j := 0; j < userInt; j++ {
		for g := 0; g < len(p.ReversePyramid[j]); g++ {
			rpd[cellId] = map[string]int{"cellId": cellId, "cellValue": p.ReversePyramid[j][g], "cellHidden": p.ReversedHiddenRows[j][g], "cellBreak": 0}
			cellId++
		}
		rpd[cellId] = map[string]int{"cellId": cellId, "cellValue": 0, "cellHidden": 0, "cellBreak": 1}
		cellId++
	}
	p.ReversePyramidData = rpd
}

func (p *Pyramid) reversePyramid(userInt int) {
	reversedPyramid := make([][]int, userInt)
	newIndex := 0
	for j := (userInt - 1); j >= 0; j-- {
		reversedPyramid[newIndex] = make([]int, len(p.Pyramid[j]))
		reversedPyramid[newIndex] = p.Pyramid[j]
		newIndex++
	}
	p.ReversePyramid = reversedPyramid
}

func (p *Pyramid) reverseHidden(userInt int) {
	rh := make([][]int, userInt)
	newIndex := 0
	for j := (userInt - 1); j >= 0; j-- {
		rh[newIndex] = make([]int, len(p.HiddenRows[j]))
		rh[newIndex] = p.HiddenRows[j]
		newIndex++
	}
	p.ReversedHiddenRows = rh
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

func (p *Pyramid) populateCells(userInt int) {
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
	p.Pyramid = pyramid
}

func (p *Pyramid) populateCellsReverse(userInt int) {
	pyramid := make([][]int, userInt)
	cntr := 0
	for x := (userInt - 1); x >= 0; x-- {
		innerLen := userInt
		pyramid[x] = make([]int, (innerLen - cntr))
		cntr++

		for j := 0; j < len(pyramid[x]); j++ {
			fmt.Println("xj", x, j)
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
	p.Pyramid2 = pyramid
}

func (p *Pyramid) determineHidden(userInt int) {
	//TODO: this actually isn't determining hidden... it's determining what to show. fix names

	pattern4 := [][]int{{0, 6, 9}, {3, 6, 8}} //TODO: this won't work actually... see notebook
	pattern5 := [][]int{{0, 1, 7, 9, 11}, {3, 4, 6, 9, 11}}

	rand.Seed(time.Now().UnixNano())
	patternToUse := rand.Intn(2)

	var pattern []int

	switch userInt {
	case 4:
		pattern = pattern4[patternToUse]
	case 5:
		pattern = pattern5[patternToUse]
	}

	hiddenRows := make([][]int, userInt)
	cntr := 0

	// if the cell is 1, it's "shown"
	for j, g := range p.Pyramid {
		hiddenRows[j] = make([]int, len(g))
		for a, _ := range g {
			for _, l := range pattern {
				if l == cntr {
					hiddenRows[j][a] = 1
				}
			}
			cntr++
		}
	}
	p.HiddenRows = hiddenRows
}
