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

func (s *StackA) Pop() {
    if s.IsEmpty() {
        return
    }
    s.items = s.items[:len(s.items)-1]
}

func (s *StackB) Push(data int) {
    s.items = append(s.items, data)
}

func (s *StackB) Pop() {
    if s.IsEmpty() {
        return
    }
    s.items = s.items[:len(s.items)-1]
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
func main()  {
	StackA := StackA{}
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
	fmt.Println(StackA.items)
}
