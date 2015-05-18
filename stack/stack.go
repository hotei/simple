// stack.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"log"
)

type StackOfIntType struct {
	name  string
	stack []int
}

func (s *StackOfIntType) SetName(name string) {
	s.name = name
}

func (s *StackOfIntType) Push(i int) {
	s.stack = append(s.stack, i)
}

func (s *StackOfIntType) Pop() int {
	slen := len(s.stack)
	if slen <= 0 {
		log.Panicf("oops too many Pops() of stack %s", s.name)
	}
	rv := s.stack[slen-1]
	s.stack = s.stack[0 : slen-1]
	return rv
}

func main() {
	fmt.Printf("Stack.go\n")

	s1 := new(StackOfIntType)

	s1.Push(10)
	s1.Push(11)
	x := s1.Pop()
	fmt.Printf("x(%d)\n", x)
	x = s1.Pop()
	fmt.Printf("x(%d)\n", x)

}
