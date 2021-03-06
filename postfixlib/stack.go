package postfixlib

import "fmt"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() interface{} {
	return s.top.value
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s Stack) String() string {
	curPtr := s.top
	str := ""
	for i := 0; i < s.size; i++ {
		str += fmt.Sprintf("%d:%v\n", i, curPtr.value)
		curPtr = curPtr.next
	}
	return str
}
