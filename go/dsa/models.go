package dsa

type ListNode[T any] struct {
	value T
	next *ListNode[T]
}

func NewListNode[T any](value T) *ListNode[T] {
	return &ListNode[T]{
		value: value,
	}
}

type TreeNode[T any] struct {
	value T
	left *TreeNode[T]
	right *TreeNode[T]
}