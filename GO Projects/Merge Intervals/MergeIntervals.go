package main

import (
    "fmt"
    "sort"
)

// Function to merge overlapping intervals
func mergeIntervals(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    // Sort intervals by the start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        last := result[len(result)-1]
        current := intervals[i]

        if current[0] <= last[1] {
            // Merge intervals
            result[len(result)-1][1] = max(last[1], current[1])
        } else {
            result = append(result, current)
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
    merged := mergeIntervals(intervals)
    fmt.Println("Merged intervals:", merged)  // Output: [[1 6] [8 10] [15 18]]
}
