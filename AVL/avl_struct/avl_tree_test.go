package avlTree

import (
	"context"
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	tree := &AvlTree{Key: 10, Height: 1, Count: 1}

	tree = tree.InsertTree(20)
	tree = tree.InsertTree(30)
	tree = tree.InsertTree(40)
	tree = tree.InsertTree(50)
	tree = tree.InsertTree(25)
	fmt.Println("AVL-tree")
	tree.Traversal()

	tree = tree.RemoveTree(30)
	fmt.Println("AVL-tree после удаления 30:")
	tree.Traversal()

	if tree.getBalanceFactor() > 1 || tree.getBalanceFactor() < -1 {
		t.Errorf("Дерево не сбалансировано после удаления 30")
	}

	if tree.Height != 3 {
		t.Errorf("Неверная высота после удаления 30, ожидали 3, имеем %d", tree.Height)
	}

	tree = tree.InsertTree(25)
	if tree.FindKey(25).Count != 2 {
		t.Errorf("Неверный count для 25, Ожидали 2, got %d", tree.Count)
	}

	tree = tree.RemoveTree(100)
	fmt.Println("Удаляем не существующий ключ 100:")
	tree.Traversal()
	ctx := context.Background()
	tree.VisualizeGraph(ctx)
}
