package main

import (
	avlTree "avl/avl_struct"
	"context"
)

func main() {
	tree := &avlTree.AvlTree{Key: 6, Height: 1, Count: 1}
	tree = tree.InsertTree(21)
	tree = tree.InsertTree(28)
	tree = tree.InsertTree(12)
	tree = tree.InsertTree(228)
	tree = tree.InsertTree(1)
	tree = tree.InsertTree(3)
	tree = tree.InsertTree(2)
	tree = tree.InsertTree(2)
	tree = tree.InsertTree(2)
	tree = tree.InsertTree(2)
	tree = tree.InsertTree(2)
	tree = tree.InsertTree(2)
	tree = tree.InsertTree(16)
	tree.Traversal()
	ctx := context.TODO()

	tree.VisualizeGraph(ctx)

}
