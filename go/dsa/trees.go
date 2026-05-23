package dsa

import "fmt"


func BuildBinaryTree[T any](values []T, idx int) *TreeNode[T] {
	if idx >= len(values) {
		return nil
	}

	node := &TreeNode[T]{
		value: values[idx],
	}

	node.left = BuildBinaryTree(values, 2 * idx + 1)
	node.right = BuildBinaryTree(values, 2 * idx + 2)

	return node
}

func PreOrder[T any](root *TreeNode[T]) {
	if root == nil {
		return 
	}

	fmt.Println(root.value)
	PreOrder(root.left)
	PreOrder(root.right)
}

func InOrder[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Println(root.value)
	InOrder(root.right)
}

func PostOrder[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.value)
}