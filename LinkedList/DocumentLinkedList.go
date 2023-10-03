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

// Size Of LL
func (l *Link) Size() int {
	return l.Count
}

//Check LL is empty
func (l *Link) IsEmpty() bool {
	return l.Count == 0
}

// Add Node into Ll
func (l *Link) AddHead(Val int) {
	l.Head = &Nod{Val: Val}
	l.Count++
}

// Add Node At last Into LL
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

//Print The linked List Value
func (l *Link) Print() {
	temp := l.Head
	for temp != nil {
		fmt.Print(temp.Val, "")
		temp = temp.Next
	}
	fmt.Println("")
}

//Searh Element in Linked List

func (l *Link) IsPresent(data int) bool {
	temp := l.Head

	for temp != nil {
		if temp.Val == data {
			return true
		}
	}
	return false
}

//Delete First Element in LL

func (l *Link) RemoveHead() (int, bool) {
	if l.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	val := l.Head.Val
	l.Head = l.Head.Next
	l.Count--
	return val, true

}
