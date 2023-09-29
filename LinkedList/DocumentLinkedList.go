package linkedlist

type Node struct {
	Next  *Node
	Value int
}

type Link struct {
	Head  *Node
	Count int
}

func (l *Link) Size() int {
	return l.Count
}
