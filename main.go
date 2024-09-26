package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackA struct {
	items []int
}

type StackB struct {
	items []int
}

func (s *StackA) Push(data int) {
	s.items = append(s.items, data)
}

func (s *StackA) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	n := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return n
}

func (s *StackB) Push(data int) {
	s.items = append(s.items, data)
}

func (s *StackB) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	n := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return n
}

func (s *StackA) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *StackB) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *StackA) sa() {
	s.items[len(s.items)-1], s.items[len(s.items)-2] = s.items[len(s.items)-2], s.items[len(s.items)-1]
}

func (s *StackA) ra() {
	n := s.Pop()
	s.items = append([]int{n}, s.items...)
}

func (s *StackA) rra() {
	n := s.items[0]
	s.items = s.items[1:]
	s.Push(n)
}

func isSorted(s []int) bool {
	min := 0
	for _, v := range s {
		if v > min {
			min = v
		} else {
			return false
		}
	}
	return true
}

func main() {
	StackA := StackA{}
	StackB := StackB{}
	num := os.Args[1]
	nums := strings.Split(num, " ")
	for _, v := range nums {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error")
			return
		}
		StackA.Push(n)
	}
	var tracker string
	for len(StackA.items) > 3 {
		tracker += "pb\n"
		n := StackA.Pop()
		StackB.Push(n)
	}
	if len(StackA.items) == 3 {
		if StackA.items[1] < StackA.items[0] && StackA.items[1] < StackA.items[2] {
			if StackA.items[0] > StackA.items[2] {
				StackA.sa()
				tracker += "sa\n"
			} else {
				StackA.ra()
				tracker += "ra\n"
			}
		} else if StackA.items[1] > StackA.items[0] && StackA.items[1] > StackA.items[2] {
			if StackA.items[0] > StackA.items[2] {
				StackA.sa()
				tracker += "sa\n"
				StackA.ra()
				tracker += "ra\n"
			} else {
				StackA.rra()
				tracker += "rra\n"
			}
		} else {
			StackA.sa()
			tracker += "sa\n"
			StackA.rra()
			tracker += "rra\n"
		}
	}
	fmt.Println("stackA: ", StackA.items)
	fmt.Println("stackB: ", StackB.items)
	fmt.Println(tracker)
}
