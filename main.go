package main

import (
	"DsaWorld/stacks"
	"fmt"
)

func main() {
	var s stacks.Stacks
	s.Push("Data Structure")
	s.Push("And")
	s.Push("Algorithms")

	fmt.Printf("%q", s)

	for !s.IsEmpty() {
		x, y := s.Pop()
		if y == true {
			fmt.Printf("%q", x)
		}
	}
}
