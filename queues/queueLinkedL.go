package queues

import "fmt"

type Node struct {
	Next *Node
	Val  int
}
type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

func (qu *Queue) Lsize() int {
	return qu.Size
}

func (qu *Queue) IzEmpty() bool {
	return qu.Size == 0
}

func (qu *Queue) Peak() (int, bool) {
	if qu.IzEmpty() {
		fmt.Println("EmaptyListError")
		return 0 , false
	}
	return qu.Head.Val,true
}

func (qu *Queue)Print(){
	temp:=qu.Head
	for temp!=nil{
		fmt.Println(temp.Val,"")
		temp = temp.Next
	}
}