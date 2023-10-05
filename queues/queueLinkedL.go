package queues


type Node struct{
	Next *Node
	Val int
}
type Queue struct{
	Head *Node
	Tail *Node
	Size int
}

func (qu *Queue)Lsize()int{
	return qu.Size
}

func (qu *Queue)IzEmpty()bool{
	return qu.Size == 0
}