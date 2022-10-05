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

func (a *LinkedList[T]) Head() *Node[T] {
	return a.head
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

type BiNode[T constraints.Ordered] struct {
	next *BiNode[T]
	Data T
	prev *BiNode[T]
}

func (a *BiNode[T]) Prev() *BiNode[T] {
	return a.prev
}

func (a *BiNode[T]) Next() *BiNode[T] {
	return a.next
}

type DoublyLinkedList[T constraints.Ordered] struct {
	head *BiNode[T]
	tail *BiNode[T]
}

func (a *DoublyLinkedList[T]) Head() *BiNode[T] {
	return a.head
}

func (a *DoublyLinkedList[T]) Tail() *BiNode[T] {
	return a.tail
}

func NewDoublyLinkedList[T constraints.Ordered]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (a *DoublyLinkedList[T]) Insert(idx uint, data T) bool {
	prev := (*BiNode[T])(nil)
	cur := a.head
	var next *BiNode[T]
	if cur != nil {
		next = cur.next
	}

	for i := uint(0); i != idx; i++ {
		if cur == nil {
			return false
		}
		prev = cur
		cur = cur.next
		if cur != nil {
			next = cur.next
		}
	}

	if prev == nil {
		a.head = &BiNode[T]{
			Data: data,
			next: cur,
			prev: nil,
		}
		if cur != nil {
			cur.prev = a.head
		}

		return true
	} else {
		new := &BiNode[T]{
			Data: data,
			next: cur,
		}
		prev.next = new
		if next != nil {
			next.prev = new
		} else {
			a.tail = new
		}
		return true
	}
}

func (a *DoublyLinkedList[T]) Append(data T) {
	if a.head == nil {
		a.head = &BiNode[T]{Data: data, next: nil, prev: nil}
		a.tail = a.head
		return
	}
	cur := a.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &BiNode[T]{Data: data, next: nil, prev: cur}
	a.tail = cur.next
}

func (a *DoublyLinkedList[T]) Find(data T) *BiNode[T] {
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

func (a *DoublyLinkedList[T]) Count(data T) (cnt int) {
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

func (a *DoublyLinkedList[T]) At(idx uint) (*BiNode[T], bool) {
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

func (a *DoublyLinkedList[T]) Set(idx uint, data T) (*BiNode[T], bool) {
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

func (a *DoublyLinkedList[T]) Length() (length int) {
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

func (a *DoublyLinkedList[T]) Delete(idx uint) bool {
	if a.head == nil {
		return false
	}

	prev := (*BiNode[T])(nil)
	cur := a.head
	var next *BiNode[T]
	if cur != nil {
		next = cur.next
	}

	for i := uint(0); i != idx; i++ {
		if cur.next == nil {
			return false
		}
		prev = cur
		cur = cur.next
		if cur != nil {
			next = cur.next
		}
	}

	if prev == nil && next == nil {
		a.head = nil
		a.tail = nil
	} else if prev != nil && next != nil {
		prev.next = cur.next
		next.prev = cur.prev
	} else if prev != nil && next == nil {
		a.tail = cur.prev
		prev.next = nil
	} else if prev == nil && next != nil {
		a.head = cur.next
		cur.prev = nil
	}

	cur = nil

	return true
}

func (a *DoublyLinkedList[T]) Remove(data T) bool {
	if a.head == nil {
		return false
	}

	prev := (*BiNode[T])(nil)
	cur := a.head
	var next *BiNode[T]
	if cur != nil {
		next = cur.next
	}
	for cur.Data != data {
		if cur.next == nil {
			return false
		}
		prev = cur
		cur = cur.next
		if cur != nil {
			next = cur.next
		}
	}

	if prev == nil && next == nil {
		a.head = nil
		a.tail = nil
	} else if prev != nil && next != nil {
		prev.next = cur.next
		next.prev = cur.prev
	} else if prev != nil && next == nil {
		a.tail = cur.prev
		prev.next = nil
	} else if prev == nil && next != nil {
		a.head = cur.next
		cur.prev = nil
	}

	cur = nil
	return true
}
