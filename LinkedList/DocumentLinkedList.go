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

func (l *Link) IsEmpty()bool{
	return l.Count == 0
}

func (l *Link)AddHead(Value int){
	curr:=l.Head
	l.Head =&Node{Value:90}
	l.Count++
}

