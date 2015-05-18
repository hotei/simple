// inventory.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// inventory vars moved to struct

package main

import (
	"fmt"
)

const MaxItem = 10

type inventorySystem struct {
	inventory   [MaxItem]int
	initialized bool
}

func (i *inventorySystem) initialize() {
	if !i.initialized {
		for j := 0; j < MaxItem; j++ {
			i.inventory[j] = j * 2
		}
		i.initialized = true
	}
}

func (i *inventorySystem) dump() {
	for j := 0; j < MaxItem; j++ {
		fmt.Printf("%d has %d in stock\n", j, i.inventory[j])
	}
}

func (i *inventorySystem) transaction(sku, howmany int) {
	if sku < 0 {
		fmt.Printf("sku[%d] has %d on hand\n", -sku, i.inventory[-sku])
	} else {
		i.inventory[sku] += howmany
	}
}

func main() {
	var stock inventorySystem
	stock.initialize()
	stock.transaction(1, 12)
	stock.dump()
	stock.transaction(-1, 0)
}
