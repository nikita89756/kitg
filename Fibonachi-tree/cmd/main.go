package main

import (
	heap "fibonachi/fibonachiHeap"
	"fmt"
)

func main() {
	fheap := heap.NewFibonacciHeap()

	fheap.Insert(5)
	fheap.Insert(3)
	fheap.Insert(8)
	fheap.Insert(1)

	fmt.Println("Minimum value:", fheap.GetMinimum())

	fmt.Println("Removing minimum value:", fheap.RemoveMinimum())

	fheap.Visualization()
}
