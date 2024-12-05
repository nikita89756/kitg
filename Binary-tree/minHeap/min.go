package minHeap

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

type MinHeap struct {
	heap []int
}

func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChild(index int) int {
	return 2*index + 1
}

func (h *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.heap[index] < h.heap[h.parent(index)] {
		h.heap[index], h.heap[h.parent(index)] = h.heap[h.parent(index)], h.heap[index]
		index = h.parent(index)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	left := h.leftChild(index)
	right := h.rightChild(index)
	smallest := index

	if left < len(h.heap) && h.heap[left] < h.heap[smallest] {
		smallest = left
	}

	if right < len(h.heap) && h.heap[right] < h.heap[smallest] {
		smallest = right
	}

	if smallest != index {
		h.heap[index], h.heap[smallest] = h.heap[smallest], h.heap[index]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) isEmpty() bool {
	return len(h.heap) == 0
}

func (h *MinHeap) size() int {
	return len(h.heap)
}

func (h *MinHeap) GetMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("Куча пуста")
	}
	return h.heap[0], nil
}

func (h *MinHeap) Insert(key int) {
	h.heap = append(h.heap, key)
	index := len(h.heap) - 1
	h.heapifyUp(index)
}

func (h *MinHeap) ExtractMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("куча пуста")
	}

	root := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown(0)
	return root, nil
}

func (h *MinHeap) DeleteNode(key int) error {
	index := -1

	for i := 0; i < len(h.heap); i++ {
		if h.heap[i] == key {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Ключ не найден в куче")
	}

	// Replace the node with the last element
	h.heap[index] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	// Restore heap property
	h.heapifyUp(index)
	h.heapifyDown(index)
	return nil
}

func (h *MinHeap) DecreaseKey(i, newValue int) error {
	if i < 0 || i >= len(h.heap) || newValue > h.heap[i] {
		return errors.New("Invalid index or new value")
	}
	h.heap[i] = newValue
	h.heapifyUp(i)
	return nil
}

func (h *MinHeap) PrintHeap() {
	for _, elem := range h.heap {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}
func (h *MinHeap) Visualize() {
	g := simple.NewDirectedGraph()

	// Create nodes
	for _, node := range h.heap {
		g.AddNode(simple.Node(node))
	}

	for i, node := range h.heap {
		left := h.leftChild(i)
		right := h.rightChild(i)
		if left < len(h.heap) {
			g.SetEdge(simple.Edge{F: simple.Node(node), T: simple.Node(h.heap[left])})
		}
		if right < len(h.heap) {
			g.SetEdge(simple.Edge{F: simple.Node(node), T: simple.Node(h.heap[right])})
		}
	}

	dotFile, err := os.Create("min_heap.dot")
	if err != nil {
		panic(err)
	}

	dotBytes, err := dot.Marshal(g, "MinHeap", "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = dotFile.Write(dotBytes)
	if err != nil {
		panic(err)
	}
	dotFile.Close()

	cmd := exec.Command("dot", "-Tpng", "-o", "min_heap.png", "min_heap.dot")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("png created")
}
