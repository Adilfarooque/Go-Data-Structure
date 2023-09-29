package linkedlist

import "fmt"

type Nod struct {
	Next *Nod
	Val  int
}

type Link struct {
	Head  *Nod
	Count int
}

func (l *Link) Size() int {
	return l.Count
}

func (l *Link) IsEmpty() bool {
	return l.Count == 0
}

func (l *Link) AddHead(Val int) {
	curr := l.Head
	l.Head = &Nod{Val: Val}
	l.Count++
}

func (l *Link) AddTail(Val int) {
	curr := l.Head
	newNod := &Nod{Val: Val}

	if curr != nil {
		l.Head = newNod
		return
	}

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNod
}

func (l *Link) Print() {
	temp := l.Head
	for temp != nil {
		fmt.Print(temp.Val,"")
		temp = temp.Next
	}
	fmt.Println("")
}
