package algo

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	got, ok := list.At(0)
	Nil(t, got)
	False(t, ok)

	got, ok = list.Set(0, 20)
	Nil(t, got)
	False(t, ok)

	list.Append(1)

	length := list.Length()
	Equal(t, 1, length)

	data, ok := list.At(0)
	True(t, ok)
	Equal(t, 1, data.Data)

	data, ok = list.Set(0, 2)
	True(t, ok)
	Equal(t, 2, data.Data)

	ok = list.Insert(0, 3)
	True(t, ok)

	ok = list.Delete(0)
	True(t, ok)

	ok = list.Remove(2)
	True(t, ok)

	ok = list.Insert(1, 3)
	False(t, ok)

	ok = list.Insert(0, 4)
	True(t, ok)
	ok = list.Insert(0, 5)
	True(t, ok)
	ok = list.Insert(0, 6)
	True(t, ok)
	ok = list.Insert(0, 7)
	True(t, ok)
	ok = list.Insert(4, 3)
	True(t, ok)

	ok = list.Insert(6, 1)
	False(t, ok)

	Equal(t, 5, list.Length())

	ok = list.Delete(0)
	True(t, ok)

	list.Delete(3)
	True(t, ok)

	Equal(t, 3, list.Length())

	list.Delete(1)
	True(t, ok)

	Equal(t, 2, list.Length())

	ok = list.Insert(1, 5)
	True(t, ok)

	ok = list.Remove(4)
	True(t, ok)
	ok = list.Remove(6)
	True(t, ok)

	Equal(t, 1, list.Length())

	found := list.Find(5)
	NotNil(t, found)

	found = list.Find(6)
	Nil(t, found)

	count := list.Count(5)
	Equal(t, 1, count)

	ok = list.Remove(5)
	True(t, ok)

	for i := 0; i < 4; i++ {
		list.Append(i + 1)
		node, _ := list.At(uint(i))
		Equal(t, i+1, node.Data)
	}

	got, ok = list.At(2393)
	Nil(t, got)
	False(t, ok)

	got, ok = list.Set(2393, 20)
	Nil(t, got)
	False(t, ok)

}
