package main

import (
	"DsaWorld/stacks"
	"fmt"
)

func main() {
	var s stacks.Stacks
	s.Push("Algorithms")
	s.Push("And")
	s.Push("Data Structure")

	fmt.Printf("%q", s)

	fmt.Println("")

	for len(s) > 0 {
		x, y := s.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
