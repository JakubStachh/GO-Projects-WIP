package main

import (
    "fmt"
    "container/heap"
)

// Define a min-heap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Function to find the kth largest element
func kthLargest(arr []int, k int) int {
    h := &MinHeap{}
    heap.Init(h)

    for i := 0; i < k; i++ {
        heap.Push(h, arr[i])
    }

    for i := k; i < len(arr); i++ {
        if arr[i] > (*h)[0] {
            heap.Pop(h)
            heap.Push(h, arr[i])
        }
    }

    return (*h)[0]
}

func main() {
    arr := []int{7, 10, 4, 3, 20, 15}
    k := 4
    fmt.Println("The 4th largest element is:", kthLargest(arr, k)) // Output: 7
}
