// sorter.go  (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"log"
	"sort"
)

const (
	BIGnumber int = 1e9
)

var (
	Rows    = 20
	Cols    = 20
	RowXCol = Rows * Cols
)

type Food2AntDist struct {
	antLoc  int
	foodLoc int
	dist    int
}

type Sequence []*Food2AntDist

func (s Sequence) Reset() Sequence {
	return make([]*Food2AntDist, 0, 100)
}

func (s Sequence) AddTo(p *Food2AntDist) Sequence {
	return append(s, p)
}

func (s Sequence) Less(i, j int) bool {
	return s[i].dist < s[j].dist
}

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Dump() {
	for _, val := range s {
		val.Dump()
	}
}

func (f *Food2AntDist) Dump() {
	fmt.Printf("from ant(%d) to food(%d) is %d units dist\n", f.antLoc, f.foodLoc, f.dist)
}

// fake the distance function for now
func manhatten_distance(a int, b int) int {
	if (a < 0) || (b < 0) {
		log.Panicf("distance < 0 is invalid, can't calculate manhatten_distance()")
	}
	return a + b
}

func dist_calc_all(m []*Food2AntDist) int {
	fmt.Printf("Calculating all distances at once\n")
	closest_dist := BIGnumber
	closest_ndx := -1
	for ndx, val := range m {
		val.dist = manhatten_distance(val.foodLoc, val.antLoc)
		if val.dist < closest_dist {
			closest_dist = val.dist
			closest_ndx = ndx
		}
	}
	return closest_ndx
}

func test_ary() {
	debug := false
	fmt.Printf("test_ary() start\n")
	var food_loc = []int{6, 2, 3, 4, 5, 12, 8}  // locations are unique
	var ant_loc = []int{31, 20, 30, 40, 50, 16} // locations are unique

	fmt.Printf("Food %v\n", food_loc)
	fmt.Printf("Ants %v\n", ant_loc)

	var food_pairs Sequence
	food_pairs = food_pairs.Reset()
	if debug {
		fmt.Printf("Food Pairs %v\n", food_pairs)
	}
	fmt.Printf("Should be nothing, food pairs is empty now\n")
	food_pairs.Dump()
	fmt.Printf("End of nothing\n")

	for _, food_val := range food_loc {
		for _, ant_val := range ant_loc {
			if debug {
				fmt.Printf("food(%d) ant(%d)\n", food_val, ant_val)
			}
			item := new(Food2AntDist)
			item.antLoc = ant_val
			item.foodLoc = food_val
			item.dist = -1
			food_pairs = food_pairs.AddTo(item)
		}
	}
	fmt.Printf("Unsorted food pairs after %d additions, no dist yet\n", len(food_pairs))
	food_pairs.Dump()

	_ = dist_calc_all(food_pairs)
	fmt.Printf("Unsorted pairs but with distance now\n")
	food_pairs.Dump()
	fmt.Printf("\n\nFood pairs sorted by distance ascending\n")
	sort.Sort(food_pairs)
	food_pairs.Dump()
}

func main() {
	fmt.Printf("test starting\n")
	test_ary()
	fmt.Printf("test ending\n")
}
