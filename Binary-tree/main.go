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

	fmt.Print("Heap: ")
	minHeap.PrintHeap()
	minHeap.Visualize()
	min, _ := minHeap.GetMin()
	fmt.Println("Минимальный элемент:", min)

	minHeap.ExtractMin()
	fmt.Print("После удаления минимума: ")
	minHeap.PrintHeap()
	minHeap.ExtractMin()
	fmt.Print("После удаления минимума: ")
	minHeap.PrintHeap()

	minHeap.DeleteNode(4)
	fmt.Print("После удаления  4: ")
	minHeap.PrintHeap()

	minHeap.DecreaseKey(2, 1)
	fmt.Print("После смены значения на 1: ")

	minHeap.PrintHeap()
	maxHeap := &maxheap.MaxHeap{}
	maxHeap.Insert(3)
	maxHeap.Insert(2)
	maxHeap.Insert(15)
	maxHeap.Insert(5)
	maxHeap.Insert(4)
	maxHeap.Insert(45)

	fmt.Print("Heap: ")
	maxHeap.PrintHeap()
	maxHeap.Visualize()
	max, _ := maxHeap.GetMax()
	fmt.Println("Максимальный элемент:", max)

	maxHeap.ExtractMax()
	fmt.Print("После удаления максимума: ")
	maxHeap.PrintHeap()
	maxHeap.ExtractMax()
	fmt.Print("После удаления максимума: ")
	maxHeap.PrintHeap()

	maxHeap.DeleteNode(4)
	fmt.Print("После удаления 4: ")
	maxHeap.PrintHeap()

	maxHeap.IncreaseKey(2, 50)
	fmt.Print("После смены значения на 50: ")
	maxHeap.PrintHeap()
}
