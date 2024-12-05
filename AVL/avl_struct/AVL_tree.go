package avlTree

import (
	"context"
	"fmt"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type AvlTree struct {
	Key    int
	Height int
	Count  int
	Left   *AvlTree
	Right  *AvlTree
}

func getHeight(tree *AvlTree) int {
	if tree == nil {
		return 0
	}
	return tree.Height
}

func (tree *AvlTree) makeRightSmallRotate() *AvlTree {
	newTree := tree.Left
	tree.Left = newTree.Right
	newTree.Right = tree
	tree.fixNodeHeight()
	newTree.fixNodeHeight()
	return newTree
}

func (tree *AvlTree) makeLeftSmallRotate() *AvlTree {
	newTree := tree.Right
	tree.Right = newTree.Left
	newTree.Left = tree
	tree.fixNodeHeight()
	newTree.fixNodeHeight()
	return newTree
}

func (tree *AvlTree) getBalanceFactor() int {
	return getHeight(tree.Right) - getHeight(tree.Left)
}

func (tree *AvlTree) fixNodeHeight() {
	leftHeight := getHeight(tree.Left)
	rightHeight := getHeight(tree.Right)
	if leftHeight > rightHeight {
		tree.Height = leftHeight + 1
	} else {
		tree.Height = rightHeight + 1
	}
}

func (tree *AvlTree) balanceTree() *AvlTree {
	tree.fixNodeHeight()
	if tree.getBalanceFactor() == 2 {
		if tree.Right.getBalanceFactor() < 0 {
			tree.Right = tree.Right.makeRightSmallRotate()
		}
		return tree.makeLeftSmallRotate()
	}
	if tree.getBalanceFactor() == -2 {
		if tree.Left.getBalanceFactor() > 0 {
			tree.Left = tree.Left.makeLeftSmallRotate()
		}
		return tree.makeRightSmallRotate()
	}
	return tree
}

func (tree *AvlTree) FindMinimalKeyInSubtree() *AvlTree {
	if tree.Left == nil {
		return tree
	}
	return tree.Left.FindMinimalKeyInSubtree()
}

func (tree *AvlTree) removeMinimalSubtree() *AvlTree {
	if tree.Left == nil {
		return tree.Right
	}
	tree.Left = tree.Left.removeMinimalSubtree()
	return tree.balanceTree()
}

func (tree *AvlTree) RemoveTree(Key int) *AvlTree {
	if Key < tree.Key {
		if tree.Left != nil {
			tree.Left = tree.Left.RemoveTree(Key)
		}
	} else if tree.Key < Key {
		if tree.Right != nil {
			tree.Right = tree.Right.RemoveTree(Key)
		}
	} else {
		minInRightSubtree := tree.Right.FindMinimalKeyInSubtree()
		minInRightSubtree.Right = tree.Right.removeMinimalSubtree()
		minInRightSubtree.Left = tree.Left
		return minInRightSubtree.balanceTree()
	}
	return tree.balanceTree()
}

func (tree *AvlTree) InsertTree(Key int) *AvlTree {
	if Key == tree.Key {
		tree.Count += 1
		return tree
	}
	if Key > tree.Key {
		if tree.Right == nil {
			tree.Right = &AvlTree{Key: Key, Height: 1, Count: 1}
		} else {
			tree.Right = tree.Right.InsertTree(Key)
		}
	} else {
		if tree.Left == nil {
			tree.Left = &AvlTree{Key: Key, Height: 1, Count: 1}
		} else {
			tree.Left = tree.Left.InsertTree(Key)
		}
	}
	return tree.balanceTree()
}

func (tree *AvlTree) FindKey(key int) *AvlTree {
	if tree == nil {
		return nil
	}

	if key == tree.Key {
		return tree
	} else if key < tree.Key {
		return tree.Left.FindKey(key)
	} else {
		return tree.Right.FindKey(key)
	}
}
func (tree *AvlTree) Traversal() {
	if tree.Left != nil {
		tree.Left.Traversal()
	}
	fmt.Println(tree.Key)
	if tree.Right != nil {
		tree.Right.Traversal()
	}
}
func (tree *AvlTree) VisualizeGraph(ctx context.Context) {
	g, err := graphviz.New(ctx)
	graph, err := g.Graph()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			panic(err)
		}
		g.Close()
	}()
	tree.Visualize(graph, nil)
	if err := g.RenderFilename(ctx, graph, graphviz.PNG, "avl_tree.png"); err != nil {
		panic(err)
	}

	fmt.Println("AVL tree visualization saved to avl_tree.png")
}
func (tree *AvlTree) Visualize(g *cgraph.Graph, parent *cgraph.Node) *cgraph.Node {
	if tree == nil {
		return nil
	}

	node, err := g.CreateNodeByName(fmt.Sprintf("%d", tree.Key))
	if err != nil {
		panic(err)
	}
	node.SetLabel(fmt.Sprintf("%d\n(h:%d),(cnt:%d)", tree.Key, tree.Height, tree.Count))

	if parent != nil {
		_, err := g.CreateEdgeByName("", parent, node)
		if err != nil {
			panic(err)
		}
	}

	tree.Left.Visualize(g, node)
	tree.Right.Visualize(g, node)

	return node
}
