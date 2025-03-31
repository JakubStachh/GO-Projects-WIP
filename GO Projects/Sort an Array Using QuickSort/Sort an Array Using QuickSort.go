package main

import "fmt"

// Function to implement QuickSort
func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    pivot := arr[0]
    left, right := 0, len(arr)-1

    for left < right {
        for arr[left] <= pivot && left < len(arr) {
            left++
        }
        for arr[right] > pivot && right > 0 {
            right--
        }
        if left < right {
            arr[left], arr[right] = arr[right], arr[left]
        }
    }

    arr[0], arr[right] = arr[right], arr[0]
    quickSort(arr[:right])
    quickSort(arr[right+1:])
    return arr
}

func main() {
    arr := []int{10, 7, 8, 9, 1, 5}
    fmt.Println("Sorted array:", quickSort(arr))  // Output: [1 5 7 8 9 10]
}
