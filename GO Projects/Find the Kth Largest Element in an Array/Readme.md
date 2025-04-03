# 📌 Kth Largest Element using Min-Heap in Go

## 🚀 Description
This Go program finds the k-th largest element in an array using a **Min-Heap**. The approach ensures efficient selection of the k-th largest element without sorting the entire array.

## 🔍 How It Works
The function `kthLargest(arr []int, k int) int`:

- Uses a **Min-Heap** (`container/heap`) to keep track of the top `k` largest elements.

- **Heap Property**: The root of the heap (smallest element) always holds the k-th largest value.

- **Heap Operations**:

  - Push the first `k` elements into the heap.

  - Iterate over the remaining elements:

    - If an element is larger than the heap root, replace it.

  - The root of the heap now holds the **k-th largest element**.

## 📜 Code Implementation
```sh
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
```
## 🎯 Example Output
```sh
The 4th largest element is: 7
```

## 📂 Explanation
### Min-Heap Implementation:
- A **Min-Heap** is used to store the `k` largest elements.

- The heap root contains the **k-th largest element**.

### Steps in `kthLargest` function:
1. Initialize a Min-Heap and push the first `k` elements.

2. Process remaining **elements**:

   - If an element is greater than the heap root (smallest in heap), replace it.

3. **Return the heap root**, which is now the k-th largest element.

### Time Complexity:
- **Heap operations** (`Push` and `Pop`) run in **O(log k)**.

- **Processing n elements** takes **O(n log k)**.

- **Overall Complexity**: **O(n log k)** (better than sorting **O(n log n)**).
