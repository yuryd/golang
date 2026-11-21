// Circular
package main

import (
	"fmt"
)

type Circular[T any] struct {
	circular []T
	used     []byte
	size     int
	maxPtr   int
	readPtr  int
	writePtr int
}

func MakeCircular[T any](size int) *Circular[T] {
	result := Circular[T]{circular: make([]T, size, size), used: make([]byte, size, size), size: size, maxPtr: size - 1, readPtr: -1}
	return &result
}

func (c *Circular[T]) Write(item T) {
	c.circular[c.writePtr] = item
	c.used[c.writePtr] = 1
	c.writePtr++
	if c.writePtr > c.maxPtr {
		c.writePtr = 0
	}
}

func (c *Circular[T]) Read() T {
	var res T
	for c.readPtr == -1 || c.used[c.readPtr] == 0 {
		return res
	}
	res = c.circular[c.readPtr]
	c.used[c.readPtr] = 1
	c.readPtr++
	if c.readPtr > c.maxPtr {
		c.readPtr = 0
	}
	return res
}

func readCirc[T any](c *Circular[T]) {
	const maxLoop int = 1000000
	var cnt int
	for cnt <= maxLoop {
		val := c.Read()
		fmt.Println(val)
		cnt++
	}
}

func main() {
	fmt.Println("Circular!")
	circ := MakeCircular[int](20)
	go readCirc(circ)
	circ.Write(15)
	circ.Write(99)
	fmt.Printf("%v\n", circ)
	ch := make(chan int)
	<-ch
	fmt.Printf("%v\n", circ)
}
