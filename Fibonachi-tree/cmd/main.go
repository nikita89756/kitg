package main

import (
	structure "fibonachi/fibonacciHeap"
	"fmt"
)

func main() {
	fheap := structure.NewFibonacciHeap()

	fheap.Insert(11)
	fheap.Insert(10)
	fheap.Insert(39)
	fheap.Insert(26)
	fheap.Insert(24)

	fmt.Printf("Minimum value: %d\n", fheap.GetMin())
	fmt.Printf("Minimum value removed: %d\n", fheap.ExtractMin())
}
