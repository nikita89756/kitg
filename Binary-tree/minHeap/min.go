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

// Get the parent index
func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

// Get the left child index
func (h *MinHeap) leftChild(index int) int {
	return 2*index + 1
}

// Get the right child index
func (h *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

// Heapify up to maintain heap property
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.heap[index] < h.heap[h.parent(index)] {
		h.heap[index], h.heap[h.parent(index)] = h.heap[h.parent(index)], h.heap[index]
		index = h.parent(index)
	}
}

// Heapify down to maintain heap property
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

// Check if the heap is empty
func (h *MinHeap) isEmpty() bool {
	return len(h.heap) == 0
}

// Get the size of the heap
func (h *MinHeap) size() int {
	return len(h.heap)
}

// Get the minimum element
func (h *MinHeap) GetMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("Heap is empty")
	}
	return h.heap[0], nil
}

// Insert a new key
func (h *MinHeap) Insert(key int) {
	h.heap = append(h.heap, key)
	index := len(h.heap) - 1
	h.heapifyUp(index)
}

// Extract the minimum element
func (h *MinHeap) ExtractMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("Heap is empty")
	}

	root := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown(0)
	return root, nil
}

// Delete a specific node
func (h *MinHeap) DeleteNode(key int) error {
	index := -1
	// Find the index of the node to delete
	for i := 0; i < len(h.heap); i++ {
		if h.heap[i] == key {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Key not found in the heap")
	}

	// Replace the node with the last element
	h.heap[index] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	// Restore heap property
	h.heapifyUp(index)
	h.heapifyDown(index)
	return nil
}

// Decrease key function
func (h *MinHeap) DecreaseKey(i, newValue int) error {
	if i < 0 || i >= len(h.heap) || newValue > h.heap[i] {
		return errors.New("Invalid index or new value")
	}
	h.heap[i] = newValue
	h.heapifyUp(i)
	return nil
}

// Print the heap elements
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

	// Create edges
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

	// Create DOT file
	dotFile, err := os.Create("min_heap.dot")
	if err != nil {
		panic(err)
	}

	// Serialize graph to DOT format
	dotBytes, err := dot.Marshal(g, "MinHeap", "", "  ")
	if err != nil {
		panic(err)
	}

	// Write DOT file
	_, err = dotFile.Write(dotBytes)
	if err != nil {
		panic(err)
	}
	dotFile.Close()

	// Convert DOT file to image using Graphviz
	cmd := exec.Command("dot", "-Tpng", "-o", "min_heap.png", "min_heap.dot")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Min heap visualization saved to min_heap.png")
}
