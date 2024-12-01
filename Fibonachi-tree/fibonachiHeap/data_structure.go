package heap

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/goccy/go-graphviz"
)

type Node struct {
	value  int
	prev   *Node
	next   *Node
	child  *Node
	parent *Node
	degree int
	marked bool
}

type FibonacciHeap struct {
	heap *Node
}

func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{heap: nil}
}

func (fh *FibonacciHeap) Insert(value int) *Node {
	newNode := createNode(value)
	fh.heap = merge(fh.heap, newNode)
	return newNode
}

func (fh *FibonacciHeap) IsEmpty() bool {
	return fh.heap == nil
}

func (fh *FibonacciHeap) Merge(other *FibonacciHeap) {
	fh.heap = merge(fh.heap, other.heap)
	other.heap = nil
}

func (fh *FibonacciHeap) DecreaseKey(n *Node, value int) {
	fh.heap = decreaseKey(fh.heap, n, value)
}

func (fh *FibonacciHeap) GetMinimum() int {
	return fh.heap.value
}

func (fh *FibonacciHeap) RemoveMinimum() int {
	old := fh.heap
	fh.heap = removeMinimum(fh.heap)
	result := old.value
	old = nil
	return result
}

func merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.value > b.value {
		a, b = b, a
	}
	aNext := a.next
	bPrev := b.prev
	a.next = b
	b.prev = a
	aNext.prev = bPrev
	bPrev.next = aNext
	return a
}

func createNode(value int) *Node {
	n := &Node{
		value:  value,
		prev:   nil,
		next:   nil,
		child:  nil,
		parent: nil,
		degree: 0,
		marked: false,
	}
	n.prev = n
	n.next = n
	return n
}

func deleteAll(n *Node) {
	if n != nil {
		current := n
		for {
			temp := current
			current = current.next
			deleteAll(temp.child)
			temp = nil
			if current == n {
				break
			}
		}
	}
}

func removeMinimum(n *Node) *Node {
	unMarkAndUnParentAll(n.child)
	if n.next == n {
		n = n.child
	} else {
		n.next.prev = n.prev
		n.prev.next = n.next
		n = merge(n.next, n.child)
	}
	if n == nil {
		return n
	}
	trees := make([]*Node, 64)
	for {
		if trees[n.degree] != nil {
			t := trees[n.degree]
			if t == n {
				break
			}
			trees[n.degree] = nil
			if n.value < t.value {
				t.prev.next = t.next
				t.next.prev = t.prev
				addChild(n, t)
			} else {
				t.prev.next = t.next
				t.next.prev = t.prev
				if n.next == n {
					t.next = t
					t.prev = t
					addChild(t, n)
					n = t
				} else {
					n.prev.next = t
					n.next.prev = t
					t.next = n.next
					t.prev = n.prev
					addChild(t, n)
					n = t
				}
			}
			continue
		} else {
			trees[n.degree] = n
		}
		n = n.next
	}
	min := n
	start := n
	for {
		if n.value < min.value {
			min = n
		}
		n = n.next
		if n == start {
			break
		}
	}
	return min
}

func unMarkAndUnParentAll(n *Node) {
	if n == nil {
		return
	}
	current := n
	for {
		current.marked = false
		current.parent = nil
		current = current.next
		if current == n {
			break
		}
	}
}

func addChild(parent, child *Node) {
	child.prev = child
	child.next = child
	child.parent = parent
	parent.degree++
	parent.child = merge(parent.child, child)
}

func decreaseKey(heap *Node, n *Node, value int) *Node {
	if n.value < value {
		return heap
	}
	n.value = value
	if n.parent != nil {
		if n.value < n.parent.value {
			heap = cut(heap, n)
			parent := n.parent
			n.parent = nil
			for parent != nil && parent.marked {
				heap = cut(heap, parent)
				n = parent
				parent = n.parent
				n.parent = nil
			}
			if parent != nil && parent.parent != nil {
				parent.marked = true
			}
		}
	} else {
		if n.value < heap.value {
			heap = n
		}
	}
	return heap
}

func cut(heap *Node, n *Node) *Node {
	if n.next == n {
		n.parent.child = nil
	} else {
		n.next.prev = n.prev
		n.prev.next = n.next
		n.parent.child = n.next
	}
	n.next = n
	n.prev = n
	n.marked = false
	return merge(heap, n)
}

func (fh *FibonacciHeap) Visualization() {
	if fh.heap == nil {
		fmt.Println("Fibonacci Heap is empty.")
		return
	}

	// Генерация DOT-кода
	dotCode := generateDotCode(fh.heap)

	// Создание файла для сохранения PNG
	pngFile, err := os.Create("fibonacci_heap.png")
	if err != nil {
		fmt.Println("Error creating PNG file:", err)
		return
	}
	defer pngFile.Close()

	// Вызов graphviz для генерации изображения
	cmd := exec.Command("dot", "-Tpng")
	cmd.Stdin = strings.NewReader(dotCode)
	cmd.Stdout = pngFile

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating visualization:", err)
		return
	}

	fmt.Println("Visualization generated: fibonacci_heap.png")
}

func generateDotCode(n *Node) string {
	var sb strings.Builder
	sb.WriteString("digraph FibonacciHeap {\n")
	sb.WriteString("  node [shape=circle];\n")

	current := n
	for {
		outputChildrenForVisualization(&sb, current)
		current = current.next
		if current == n {
			break
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}

func outputChildrenForVisualization(sb *strings.Builder, n *Node) {
	sb.WriteString(fmt.Sprintf("  \"%p\" [label=\"%d\"];\n", n, n.value))
	if n.child != nil {
		current := n.child
		for {
			sb.WriteString(fmt.Sprintf("  \"%p\" -> \"%p\";\n", n, current))
			outputChildrenForVisualization(sb, current)
			current = current.next
			if current == n.child {
				break
			}
		}
	}
}
