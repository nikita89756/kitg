package main

import (
	maxheap "bt/maxHeap"
	"bt/minHeap"
	"fmt"
)

func main() {
	minHeap := &minHeap.MinHeap{}
	minHeap.Insert(3)
	minHeap.Insert(2)
	minHeap.Insert(15)
	minHeap.Insert(5)
	minHeap.Insert(4)
	minHeap.Insert(45)

	fmt.Print("Min Heap: ")
	minHeap.PrintHeap()
	minHeap.Visualize()
	min, _ := minHeap.GetMin()
	fmt.Println("Minimum element:", min)

	minHeap.ExtractMin()
	fmt.Print("After extracting min: ")
	minHeap.PrintHeap()
	minHeap.ExtractMin()
	fmt.Print("After extracting min: ")
	minHeap.PrintHeap()

	// Delete node from min-heap
	minHeap.DeleteNode(4)
	fmt.Print("After deleting 4: ")
	minHeap.PrintHeap()

	// Decrease key in min-heap
	minHeap.DecreaseKey(2, 1)
	fmt.Print("After decreasing key at index 2 to 1: ")

	minHeap.PrintHeap()
	maxHeap := &maxheap.MaxHeap{}
	maxHeap.Insert(3)
	maxHeap.Insert(2)
	maxHeap.Insert(15)
	maxHeap.Insert(5)
	maxHeap.Insert(4)
	maxHeap.Insert(45)

	fmt.Print("Max Heap: ")
	maxHeap.PrintHeap()
	maxHeap.Visualize()
	max, _ := maxHeap.GetMax()
	fmt.Println("Maximum element:", max)

	maxHeap.ExtractMax()
	fmt.Print("After extracting max: ")
	maxHeap.PrintHeap()
	maxHeap.ExtractMax()
	fmt.Print("After extracting max: ")
	maxHeap.PrintHeap()

	// Delete node from max-heap
	maxHeap.DeleteNode(4)
	fmt.Print("After deleting 4: ")
	maxHeap.PrintHeap()

	// Increase key in max-heap
	maxHeap.IncreaseKey(2, 50)
	fmt.Print("After increasing key at index 2 to 50: ")
	maxHeap.PrintHeap()
}
