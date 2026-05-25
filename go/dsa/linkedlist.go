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

func HasCycle[T any](head *ListNode[T]) *ListNode[T] {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			slow = head

			for slow != fast {
				slow = slow.next
				fast = fast.next
			}

			return slow
		}
	}

	return nil
}