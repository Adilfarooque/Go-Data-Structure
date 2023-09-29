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
	mylist := linkedlist.LinkedList{}
	nord1 := &linkedlist.Node{Data: 90}
	nord2 := &linkedlist.Node{Data: 243}
	nord3 := &linkedlist.Node{Data: 23}
	nord4 := &linkedlist.Node{Data: 73}
	nord5 := &linkedlist.Node{Data: 202}
	nord6 := &linkedlist.Node{Data: 76}
	nord7 := &linkedlist.Node{Data: 22}
	nord8 := &linkedlist.Node{Data: 987}
	nord9 := &linkedlist.Node{Data: 276}
	mylist.Prepend(nord1)
	mylist.Prepend(nord2)
	mylist.Prepend(nord3)
	mylist.Prepend(nord4)
	mylist.Prepend(nord5)
	mylist.Prepend(nord6)
	mylist.Prepend(nord7)
	mylist.Prepend(nord8)
	mylist.Prepend(nord9)

	mylist.PrintListData()

	//Single Linked List Documentation
	mList:=linkedlist.Link{}
	mList.Head()

}
