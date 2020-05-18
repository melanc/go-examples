package stack

import (
	"fmt"
	"testing"
)

// 切片实现栈
type Stack struct {
	data []int
}

func TestStack(t *testing.T) {
	var a []int = []int{6, 1, 2, 5, 9, 3, 4, 7, 10, 8}
	s := NewStack(a)
	PrintStack(s)
	Push(s, 11)
	Push(s, 13)
	PrintStack(s)
	Pop(s)
	PrintStack(s)
}

func NewStack(arr []int) *Stack {
	return &Stack{arr}
}

func Size(s *Stack) int {
	return len(s.data)
}

func Push(s *Stack, val int) {
	s.data = append(s.data, val)
}

func Pop(s *Stack) int {
	l := Size(s)
	val := s.data[l-1]
	s.data = s.data[:l-1]
	return val
}

func PrintStack(s *Stack) {
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Printf("%v ", s.data[i])
	}
	fmt.Println("")
}
