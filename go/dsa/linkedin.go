package dsa

import "fmt"

func BuildLinkedList[T any](values []T) *ListNode[T] {
	dummy := &ListNode[T]{}
	tail := dummy

	for _, value := range values {
		node := NewListNode(value)
		tail.next = node
		tail = tail.next
	}

	return dummy.next
}

func TraverseLinkedList[T any](head *ListNode[T]) {
	curr := head

	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}