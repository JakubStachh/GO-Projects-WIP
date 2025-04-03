package main

import "fmt"

// Function to merge two sorted arrays
func mergeArrays(arr1, arr2 []int) []int {
    result := []int{}
    i, j := 0, 0

    for i < len(arr1) && j < len(arr2) {
        if arr1[i] < arr2[j] {
            result = append(result, arr1[i])
            i++
        } else {
            result = append(result, arr2[j])
            j++
        }
    }

    // Add remaining elements
    for i < len(arr1) {
        result = append(result, arr1[i])
        i++
    }
    for j < len(arr2) {
        result = append(result, arr2[j])
        j++
    }

    return result
}

func main() {
    arr1 := []int{1, 3, 5, 7}
    arr2 := []int{2, 4, 6, 8}
    fmt.Println(mergeArrays(arr1, arr2)) // Output: [1 2 3 4 5 6 7 8]
}
