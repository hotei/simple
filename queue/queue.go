// queue.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"log"
	"sort"
)

type QueueOfIntType struct {
	Name     string
	Queue    []int
	LenQueue int
}

// ============================================================

func (q *QueueOfIntType) SetName(name string) {
	q.Name = name
}

func (q *QueueOfIntType) Dump() {
	fmt.Printf("Queue '%s' Dump([%d])\n", q.Name, q.LenQueue)
	for ndx, val := range q.Queue {
		fmt.Printf("[%d] = %d\n", ndx, val)
	}
}

func (q *QueueOfIntType) Sort_Asc() {
	sort.Sort(q)
}

func (q *QueueOfIntType) Less(i, j int) bool {
	return q.Queue[i] >= q.Queue[j]
}

func (q *QueueOfIntType) Swap(i, j int) {
	q.Queue[i], q.Queue[j] = q.Queue[j], q.Queue[i]
}

func (q *QueueOfIntType) Len() int {
	return len(q.Queue)
}

func (q *QueueOfIntType) AddFrnt(n int) {
	fmt.Printf("Adding %d to front\n", n)
	tmp := make([]int, q.LenQueue+1)
	tmp[0] = n
	for ndx, val := range q.Queue {
		tmp[ndx+1] = val
	}
	q.Queue = tmp
	q.LenQueue++
	return
}

func (q *QueueOfIntType) AddTail(n int) {
	fmt.Printf("Adding %d to tail\n", n)
	q.Queue = append(q.Queue, n)
	q.LenQueue++
	return
}

func (q *QueueOfIntType) PopFrnt() int {
	rc := q.Queue[0]
	q.Queue = q.Queue[1:]
	q.LenQueue--
	return rc
}

func (q *QueueOfIntType) PopTail() int {
	rc := q.Queue[q.LenQueue-1]
	q.Queue = q.Queue[:q.LenQueue-1]
	q.LenQueue--
	return rc
}

func (q *QueueOfIntType) PeekFrnt() int {
	return q.Queue[0]
}

func (q *QueueOfIntType) PeekTail() int {
	return q.Queue[q.LenQueue-1]
}

func (q *QueueOfIntType) Reset() {
	q.LenQueue = 0
	q.Name = ""
	q.Queue = make([]int, 0)
}

func main() {
	fmt.Printf("Queue.go\n")

	q1 := new(QueueOfIntType)
	q1.SetName("sorted numbers")
	q1.AddTail(10)
	q1.Dump()
	q1.AddTail(11)
	q1.Dump()
	q1.AddFrnt(9)
	q1.Dump()
	q1.AddFrnt(8)
	q1.Dump()
	q1.AddFrnt(1)
	q1.Dump()

	fmt.Printf("Sort Ascending\n")
	q1.Sort_Asc()
	q1.Dump()
	var x int

	x = q1.PopTail()
	fmt.Printf("PopTail x(%d)\n", x)
	q1.Dump()

	x = q1.PeekFrnt()
	fmt.Printf("PeekFrnt x(%d)\n", x)
	q1.Dump()

	x = q1.PopFrnt()
	fmt.Printf("PopFrnt x(%d)\n", x)
	q1.Dump()

	x = q1.PeekTail()
	fmt.Printf("PeekTail x(%d)\n", x)
	q1.Dump()

	fmt.Printf("Reset()\n")
	q1.Reset()
	q1.Dump()

	if false {
		log.Panicf("don't panic yet")
	}
}
