package algo

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Data T
	next *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

type LinkedList[T constraints.Ordered] struct {
	head *Node[T]
}

func NewLinkedList[T constraints.Ordered]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
	}
}

func (a *LinkedList[T]) Insert(idx uint, data T) bool {
	prev := (*Node[T])(nil)
	cur := a.head

	for i := uint(0); i != idx; i++ {
		if cur == nil {
			return false
		}
		prev = cur
		cur = cur.next
	}

	if prev == nil {
		a.head = &Node[T]{
			Data: data,
			next: cur,
		}
		return true
	} else {
		prev.next = &Node[T]{
			Data: data,
			next: cur,
		}
		return true
	}
}

func (a *LinkedList[T]) Append(data T) {
	if a.head == nil {
		a.head = &Node[T]{Data: data, next: nil}
		return
	}
	cur := a.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &Node[T]{Data: data, next: nil}
}

func (a *LinkedList[T]) Find(data T) *Node[T] {
	if a.head == nil {
		return nil
	}

	cur := a.head
	for cur.Data != data {
		if cur.next == nil {
			return nil
		}
		cur = cur.next
	}

	return cur
}

func (a *LinkedList[T]) Count(data T) (cnt int) {
	if a.head == nil {
		return
	}

	cur := a.head
	for cur != nil {
		if cur.Data == data {
			cnt++
		}
		cur = cur.next
	}

	return
}

func (a *LinkedList[T]) At(idx uint) (*Node[T], bool) {
	if a.head == nil {
		return nil, false
	}

	cur := a.head
	for i := uint(0); i != idx; i++ {
		if cur.next == nil {
			return nil, false
		}
		cur = cur.next
	}

	return cur, true
}

func (a *LinkedList[T]) Set(idx uint, data T) (*Node[T], bool) {
	if a.head == nil {
		return nil, false
	}

	cur := a.head
	for i := uint(0); i != idx; i++ {
		if cur.next == nil {
			return nil, false
		}
		cur = cur.next
	}

	cur.Data = data
	return cur, true
}

func (a *LinkedList[T]) Length() (length int) {
	if a.head == nil {
		return
	}

	cur := a.head
	length = 1
	for cur.next != nil {
		length++
		cur = cur.next
	}

	return
}

func (a *LinkedList[T]) Delete(idx uint) bool {
	if a.head == nil {
		return false
	}

	prev := (*Node[T])(nil)
	cur := a.head
	for i := uint(0); i != idx; i++ {
		if cur.next == nil {
			return false
		}
		prev = cur
		cur = cur.next
	}

	if prev != nil {
		prev.next = cur.next
	} else {
		a.head = cur.next
	}

	cur = nil

	return true
}

func (a *LinkedList[T]) Remove(data T) bool {
	if a.head == nil {
		return false
	}

	prev := (*Node[T])(nil)
	cur := a.head
	for cur.Data != data {
		if cur.next == nil {
			return false
		}
		prev = cur
		cur = cur.next
	}

	if prev != nil {
		prev.next = cur.next
	} else {
		a.head = cur.next
	}
	cur = nil

	return true
}
