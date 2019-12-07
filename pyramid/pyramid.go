package pyramid

//TODO: Generate documentation

import (
	"math/rand"
	"time"
)

type Pyramid struct {
	Pyramid    [][]int
	HiddenRows [][]int
}

func (p *Pyramid) GeneratePyramid(userInt int) *Pyramid {
	/*
		[0] = [3, 5, 10, 10]
		[1] = [8, 15, 20]
		[2] = [23, 35]
		[3] = [58]
	*/
	p.populateCells(userInt)
	p.determineHidden(userInt)
	return p
}

func (p *Pyramid) PyramidSize() int {
	var cells int = 0
	for x := 0; x < len((p.Pyramid)); x++ {
		// de-reference the pyramid, then get the slice
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

func (p *Pyramid) determineHidden(userInt int) {
	//TODO: this actually isn't determining hidden... it's determining what to show. fix names

	pattern4 := [][]int{{0, 6, 9}, {3, 6, 8}} //TODO: this won't work actually... see notebook
	pattern5 := [][]int{{0, 1, 7, 9}, {3, 4, 6, 11}}

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
