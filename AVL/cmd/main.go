package main

import (
	avlTree "avl/avl_struct"
	"context"
	"fmt"
)

func main() {
	tree := &avlTree.AvlTree{Key: 6, Height: 1, Count: 1}
	tree = tree.InsertTree(21)
	tree = tree.InsertTree(28)
	tree = tree.InsertTree(12)
	tree = tree.InsertTree(8)
	tree = tree.InsertTree(15)
	fmt.Printf("min:%d \n", tree.FindMinimalKeyInSubtree().Key)
	tree.Traversal()
	ctx := context.TODO()

	tree.VisualizeGraph(ctx)

}
