package linkedlist

type LNode struct {
	next *LNode
	val  int
}

type LLink struct {
	head1  *LNode
	count1 int
}

func (ll *LLink) IsEmpty() bool {
	return ll.count1 == 0
}

func (ll *LLink) SortedList(val int) {
	newNode := &LNode{val: val}
	curr := ll.head1

	if curr == nil || curr.val > val {
		newNode.next = ll.head1
		ll.head1 = newNode
		return
	}
	for curr.next != nil && curr.next.val < val {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode

}
