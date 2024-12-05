package structure

import (
	"fmt"
	"math"
)

type FibonacciTree struct {
	key      int
	children []*FibonacciTree
	parent   *FibonacciTree
	marked   bool
	order    int
}

func NewFibonacciTree(key int) *FibonacciTree {
	return &FibonacciTree{
		key:      key,
		children: []*FibonacciTree{},
		parent:   nil,
		marked:   false,
		order:    0,
	}
}

func (ft *FibonacciTree) AddAtEnd(t *FibonacciTree) {
	ft.children = append(ft.children, t)
	t.parent = ft
	ft.order++
}

type FibonacciHeap struct {
	trees []*FibonacciTree
	least *FibonacciTree
	count int
}

func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{
		trees: []*FibonacciTree{},
		least: nil,
		count: 0,
	}
}

func (fh *FibonacciHeap) Insert(key int) {
	newTree := NewFibonacciTree(key)
	fh.trees = append(fh.trees, newTree)
	if fh.least == nil || key < fh.least.key {
		fh.least = newTree
	}
	fh.count++
}

func (fh *FibonacciHeap) GetMin() int {
	if fh.least == nil {
		return -1
	}
	return fh.least.key
}

func (fh *FibonacciHeap) ExtractMin() int {
	smallest := fh.least
	if smallest != nil {
		for _, child := range smallest.children {
			child.parent = nil
			fh.trees = append(fh.trees, child)
		}
		fh.trees = remove(fh.trees, smallest)
		if len(fh.trees) == 0 {
			fh.least = nil
		} else {
			fh.least = fh.trees[0]
			fh.consolidate()
		}
		fh.count--
		return smallest.key
	}
	return -1
}

func (fh *FibonacciHeap) consolidate() {
	aux := make([]*FibonacciTree, floorLog2(fh.count)+1)
	for len(fh.trees) > 0 {
		x := fh.trees[0]
		order := x.order
		fh.trees = remove(fh.trees, x)
		for aux[order] != nil {
			y := aux[order]
			if x.key > y.key {
				x, y = y, x
			}
			x.AddAtEnd(y)
			aux[order] = nil
			order++
		}
		aux[order] = x
	}
	fh.least = nil
	for _, k := range aux {
		if k != nil {
			fh.trees = append(fh.trees, k)
			if fh.least == nil || k.key < fh.least.key {
				fh.least = k
			}
		}
	}
}

func (fh *FibonacciHeap) DecreaseKey(x *FibonacciTree, newKey int) error {
	if newKey > x.key {
		return fmt.Errorf("new key is greater than current key")
	}
	x.key = newKey
	y := x.parent
	if y != nil && x.key < y.key {
		fh.cut(x, y)
		fh.cascadingCut(y)
	}
	if x.key < fh.least.key {
		fh.least = x
	}
	return nil
}

func (fh *FibonacciHeap) cut(x, y *FibonacciTree) {
	y.children = remove(y.children, x)
	y.order--
	x.parent = nil
	x.marked = false
	fh.trees = append(fh.trees, x)
}

func (fh *FibonacciHeap) cascadingCut(y *FibonacciTree) {
	z := y.parent
	if z != nil {
		if !y.marked {
			y.marked = true
		} else {
			fh.cut(y, z)
			fh.cascadingCut(z)
		}
	}
}

func (fh *FibonacciHeap) Delete(x *FibonacciTree) {
	fh.DecreaseKey(x, math.MinInt32)
	fh.ExtractMin()
}

func floorLog2(x int) int {
	return int(math.Log2(float64(x)))
}

func remove(slice []*FibonacciTree, item *FibonacciTree) []*FibonacciTree {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
