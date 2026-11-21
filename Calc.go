// Stack
package main

import (
	"fmt"
)

type Numeric interface{ int | float32 | float64 }

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
	s.size++
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
	s.size--
	return res
}

func (s *Stack[T]) Read() T {
	if s.ptr == -1 {
		var dummy T
		return dummy
	}
	return s.stack[s.ptr]
}

func (s *Stack[T]) Len() int {
	return s.size
}

type Calc[T Numeric] struct {
	stack *Stack[T]
}

func (c *Calc[T]) Push(val T) {
	c.stack.Push(val)
}

func (c *Calc[T]) Push2(val1, val2 T) {
	c.stack.Push(val1)
	c.stack.Push(val2)
}

func (c *Calc[T]) Add() {
	val1, val2 := c.Get2()
	c.stack.Push(val1 + val2)
}

func (c *Calc[T]) Sub() {
	val1, val2 := c.Get2()
	c.stack.Push(val1 - val2)
}

func (c *Calc[T]) SubR() {
	val1, val2 := c.Get2()
	c.stack.Push(val2 - val1)
}

func (c *Calc[T]) Sum() T {
	var sum T
	for c.stack.Len() > 0 {
		sum += c.stack.Pop()
	}
	return sum
}

func (c *Calc[T]) Avg() float64 {
	len := c.stack.Len()
	sum := c.Sum()
	return sum / len
}

func (c *Calc[T]) Get() T {
	return c.stack.Pop()
}

func (c *Calc[T]) Get2() (T, T) {
	val2 := c.stack.Pop()
	val1 := c.stack.Pop()
	return val1, val2
}

func (c *Calc[T]) Read() T {
	return c.stack.Read()
}

func MakeCalc[T Numeric](size int) Calc[T] {
	result := Calc[T]{stack: MakeStack[T](size)}
	return result
}

func main() {
	fmt.Println("Calc!")
	c := MakeCalc[int](50)
	c.Push(15)
	c.Push(25)
	c.Add()
	fmt.Println(c.Read(), c.Get(), c.Read())
	c.Push2(15, 25)
	c.Sub()
	c.Push(20)
	c.Add()
	fmt.Println(c.Read(), c.Get(), c.Read())
}
