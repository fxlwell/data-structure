package data_struct

import (
	"fmt"
)

type Stack struct {
	Top  int
	Data []int
}

func NewStack() *Stack {
	s := &Stack{
		Top:  -1,
		Data: make([]int, 100),
	}
	return s
}

func (s *Stack) Push(d int) {
	s.Top++
	s.Data[s.Top] = d
}

func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("is empty")
	}
	v := s.Data[s.Top]
	s.Top--
	return v, nil
}
