package linkedlist

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (L *LinkedList) Prepend(n *Node) {
	Second := L.Head
	L.Head = n
	L.Head.Next = Second
	L.Length++
}
