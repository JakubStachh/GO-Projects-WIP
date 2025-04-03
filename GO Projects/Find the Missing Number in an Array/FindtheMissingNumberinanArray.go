package main

import "fmt"

// Function to find missing number
func findMissingNumber(arr []int, n int) int {
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    for _, num := range arr {
        actualSum += num
    }
    return expectedSum - actualSum
}

func main() {
    arr := []int{1, 2, 4, 6, 3, 7, 8}
    n := 8
    fmt.Println("Missing number:", findMissingNumber(arr, n)) // Output: 5
}
