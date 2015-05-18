// map.go

/*
 * NB: if key value not present in map demo below,  int values default to zero
 * to determine if a key is present see the "comma ok" example
 */

package main

import (
	"fmt"
	"sort"
)

type STM map[string]int

func (m STM) Len() int {
	return len(m)
}

func (m STM) Less(a, b string) bool {
	return m[a] < m[b]
}

func map_test_one() {
	var exists bool
	m := STM{"one": 100, "two": 22, "four": 44}

	rc := m["one"] // note that here we don't check if value exists or not
	fmt.Printf(`m["one"] = %v`+"\n", rc)

	rc = m["two"]
	fmt.Printf(`m["two"] = %v`+"\n", rc)

	rc = m["four"]
	fmt.Printf(`m["four"] = %v`+"\n", rc)

	rc, exists = m["three"]
	fmt.Printf(`m["three"] (unmapped value) = %v, exists = %v`+"\n", rc, exists)

	fmt.Printf("assign m[\"three\"] = 333\n")
	m["three"] = 333
	rc, exists = m["three"]
	fmt.Printf("m[\"three\"] (mapped value) = %v, exists = %v\n", rc, exists)

	fmt.Printf("Length of map = %d\n", len(m))
	for ndx, each := range m {
		fmt.Printf("%s %d\n", ndx, each)
	}
	n := len(m)
	// there was a breaking syntax change, this fails to remove element now
	fmt.Printf(`Remove m["two"]` + "\n")
	m["two"], exists = 0, false
	fmt.Printf("Length of map = %d\n", len(m))
	if n == len(m) {
		fmt.Printf("Length of map didn't change => failure (expected)\n")
	}
	fmt.Printf("Use the delete keyword instead to remove a map member\n")
	delete(m, "two") // no map or no element is no problem
	fmt.Printf("Length of map = %d\n", len(m))
	if len(m) == (n - 1) {
		fmt.Printf("delete of element succeeded\n")
	}
	for ndx, each := range m {
		fmt.Printf("%s %d\n", ndx, each)
	}

	// You can do something like this in python, below is go equivalent
	//  keys = dict.Keys()
	// 	keys.Sort()
	//	for each in keys {
	//		fmt.Printf("%s %d\n", each, dict[each])
	//	}

	fmt.Printf("Sorting by map key example \n")
	keys := make([]string, 0, 30)
	for ndx, _ := range m {
		keys = append(keys, ndx)
	}
	fmt.Printf("Un-sorted keys = %v\n", keys)
	sort.Strings(keys)
	fmt.Printf("Alpha sorted keys = %v\n", keys)
	fmt.Printf("Length of map = %d\n", len(m))
	for ndx, each := range keys {
		fmt.Printf("key(%s) value(%d)\n", keys[ndx], m[each])
	}

}

func main() {
	map_test_one()
}
