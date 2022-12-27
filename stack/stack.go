package stack

/*
* @author: Chen Chiheng
* @date: 2022/12/27 0027 10:21
* @version: 1.0
* @description:
**/

type StackInterface interface {
	Push(val interface{})
	Pop() interface{}
	Len() int
	IsEmpty() bool
}

type Stack struct {
	data     []interface{}
	topIndex int
}

func New() *Stack {
	return &Stack{data: make([]interface{}, 8, 8), topIndex: -1}
}

func (s *Stack) Push(val interface{}) {
	s.topIndex++
	if s.topIndex >= len(s.data) {
		s.data = append(s.data, val)
	} else {
		s.data[s.topIndex] = val
	}
}

func (s *Stack) IsEmpty() bool {
	return s.topIndex == -1
}

func (s *Stack) Len() int {
	if s.topIndex == -1 {
		return 0
	} else {
		return s.topIndex + 1
	}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	val := s.data[s.topIndex]
	s.topIndex--
	return val
}
