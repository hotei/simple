// twoDArray.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// written for AI Ant Challenge

package main

import (
	"fmt"
	"log"
)

const (
	MapRows = 5
	MapCols = 10
)

var twoDarray = [2][2]int{{0, 1}, {2, 3}}

func locationFromRowCol(row int, col int) int {
	return row*MapRows + col
}

func main() {
	fmt.Printf("twoDArray started\n")
	log.SetFlags(log.Lshortfile)
	Territory := make([]int, MapRows*MapCols)
	fmt.Printf("uncharted Territory %v\n", Territory)
	tcount := 0
	for i := 0; i < MapRows; i++ {
		for j := 0; j < MapCols; j++ {
			location := locationFromRowCol(i, j)
			Territory[location] = '?'
			tcount += 1
		}
	}
	fmt.Printf("%d in charted Territory %v\n", tcount, Territory)
	if false {
		fmt.Printf("Test\n")
		log.Panic("Panic on purpose to inspect code")
	}
	fmt.Printf("twoDArray finished\n")
}
