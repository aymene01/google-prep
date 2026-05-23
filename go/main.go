package main

import (
	"fmt"

	"github.com/aymene01/google-prep/go/dsa"
)

func main() {
	delimiter("Hello world")
	fmt.Println("helllo gooooogle")

	delimiter("LinkedLists")
	ll := dsa.BuildLinkedList([]int{1, 2, 3, 4})
	dsa.TraverseLinkedList(ll)

	delimiter("Trees")
	tree := dsa.BuildBinaryTree([]int{1, 3, 4, 5}, 0)
	dsa.PostOrder(tree)
	dsa.PreOrder(tree)
	dsa.InOrder(tree)
}


func delimiter(element string) {
	fmt.Println()
	fmt.Println(element, "-----------------------------------------------------------------------", element)
	fmt.Println()
}