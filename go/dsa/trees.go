package dsa

import (
	"fmt"
)


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

func TreeBfs[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}

	queue := []*TreeNode[T]{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func TreeDfsRecursive[T any](root *TreeNode[T]) []T {
	if root == nil {
		return []T{}
	}

	left := TreeDfsRecursive(root.left)
	right := TreeDfsRecursive(root.right)

	result := []T{root.value}
	
	result = append(result, left...)
	result = append(result, right...)

	return result
}