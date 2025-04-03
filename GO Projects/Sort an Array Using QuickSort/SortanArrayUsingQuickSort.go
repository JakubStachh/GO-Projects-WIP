package main

import "fmt"

// Function to implement QuickSort
func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    pivot := arr[0]
    left, right := 0, len(arr)-1

    // Partitioning step
    for left < right {
        for left < len(arr) && arr[left] <= pivot {
            left++
        }
        for right >= 0 && arr[right] > pivot {
            right--
        }
        if left < right {
            arr[left], arr[right] = arr[right], arr[left]
        }
    }

    arr[0], arr[right] = arr[right], arr[0] // Place pivot in the correct position

    // Recursively sort the subarrays
    quickSort(arr[:right])   // Sort left subarray
    quickSort(arr[right+1:]) // Sort right subarray

    return arr
}

func main() {
    arr := []int{10, 7, 8, 9, 1, 5}
    fmt.Println("Sorted array:", quickSort(arr))  // Output: [1 5 7 8 9 10]
}
