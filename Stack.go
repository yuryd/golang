// Stack
package main

import (
	"fmt"
)

type Stack[T any] struct {
	stack  []T
	ptr    int
	size   int
	maxPtr int
}

func MakeStack[T any](size int) *Stack[T] {
	result := Stack[T]{stack: make([]T, size, size), ptr: -1, size: size, maxPtr: size - 1}
	return &result
}

func (s *Stack[T]) Push(val T) {
	if s.ptr == s.maxPtr {
		return
	}
	s.ptr++
	s.stack[s.ptr] = val
}

func (s *Stack[T]) Pop() T {
	if s.ptr == -1 {
		var dummy T
		return dummy
	}
	var dummy T
	res := s.stack[s.ptr]
	s.stack[s.ptr] = dummy
	s.ptr--
	return res
}

func (s *Stack[T]) Len() int {
	return s.ptr + 1
}

func main() {
	fmt.Println("Stack!")
	s := MakeStack[int](1000)
	s.Push(15)
	s.Push(38)
	fmt.Printf("%d %v\n", s.Len(), s)
	s.Pop()
	fmt.Printf("%d %v\n", s.Len(), s)
}
