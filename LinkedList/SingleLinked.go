package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Leng int
}

func (l *LinkedList) Prepend(val *Node) {
	Sval := l.Head
	l.Head = val
	l.Head.Next = Sval
}

func (l LinkedList) PrintListData() {
	Toprint := l.Head
	for l.Leng != 0 {
		fmt.Printf("%d", Toprint.Data)
		Toprint = Toprint.Next
		l.Leng--
	}
	fmt.Println("")
}
