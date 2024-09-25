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
	var tracker []string
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
	for len(StackA.items) > 3 {
		tracker = append(tracker, "pb")
		n := StackA.Pop()
		StackB.Push(n)
	}
	fmt.Println("stackA: ",StackA.items)
	fmt.Println("stackB: ",StackB.items)
	fmt.Println(tracker)
}
