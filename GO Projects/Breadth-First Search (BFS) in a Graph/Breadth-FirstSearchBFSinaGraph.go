package main

import "fmt"

// Function to perform BFS
func bfs(graph map[int][]int, start int) {
    visited := make(map[int]bool)
    queue := []int{start}
    visited[start] = true

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        fmt.Print(node, " ")

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}

func main() {
    // Representing the graph as an adjacency list
    graph := map[int][]int{
        0: {1, 2},
        1: {0, 3, 4},
        2: {0, 5},
        3: {1},
        4: {1},
        5: {2},
    }

    fmt.Print("BFS Traversal: ")
    bfs(graph, 0)  // Output: BFS Traversal: 0 1 2 3 4 5
}
