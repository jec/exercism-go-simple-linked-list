package linkedlist

import "fmt"

type List struct {
	head *Element
}

type Element struct {
	songId int
	next   *Element
}

// New creates and returns a List containing the values in `elements`
func New(elements []int) *List {
	list := List{}

	for _, songId := range elements {
		list.Push(songId)
	}

	return &list
}

// Size returns the number of songIds in the List
func (list *List) Size() int {
	count := 0
	for ptr := list.head; ptr != nil; ptr = ptr.next {
		count += 1
	}
	return count
}

// Push inserts an element at the head of List
func (list *List) Push(element int) {
	list.head = &Element{songId: element, next: list.head}
}

// Pop removes the head element from the List and returns its songId
func (list *List) Pop() (int, error) {
	head := list.head
	if head == nil {
		return 0, fmt.Errorf("list is empty")
	}

	list.head = head.next
	return head.songId, nil
}

// Array returns an array of songIds from the List
func (list *List) Array() []int {
	size := list.Size()
	result := make([]int, size)
	for i, ptr := size-1, list.head; ptr != nil; i, ptr = i-1, ptr.next {
		result[i] = ptr.songId
	}
	return result
}

// Reverse returns a List with the songIds in reverse order
func (list *List) Reverse() *List {
	newList := List{}
	for ptr := list.head; ptr != nil; ptr = ptr.next {
		newList.Push(ptr.songId)
	}
	return &newList
}
