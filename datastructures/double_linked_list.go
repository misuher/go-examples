package doubleLinkedList

type List struct {
	head, tail *Node
}

type Node struct {
	value      string
	next, prev *Node
}

func (l *List) First() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Push(val string) *List {
	n := &Node{value: val}
	if l.head == nil { //first node
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}

	l.tail = n
	return l
}

func (l *List) Pop() string {
	if l.tail == nil {
		return ""
	}
	value := l.tail.value
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	}
	return value
}

func (l *List) Find(val string) *Node {
	for n := l.First(); n != nil; n = n.Next() {
		if n.value == val {
			return n
		}
	}
	return nil
}

func (l *List) Erase(val string) bool {
	node := l.Find(val)
	if node != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
		return true
	}
	return false
}
