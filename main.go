package main

import (
	linkedlist "DsaWorld/LinkedList"
	"DsaWorld/stacks"
	"fmt"

)

func main() {
	//Reversed a String
	ipt := "Adil Farooque"
	result := stacks.ReversedString(ipt)
	fmt.Println(result)

	//isvalid Question Neetcode.
	input := "([]{})" // Replace with your input string
	if stacks.IsValid(input) {
		fmt.Println("Valid parentheses")
	} else {
		fmt.Println("Invalid parentheses")
	}

	// Stack slice
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

	//Single Linked List
	

}
