package main

import "fmt"

// Function to perform DFS
func dfs(graph map[int][]int, node int, visited map[int]bool) {
    visited[node] = true
    fmt.Print(node, " ")

    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            dfs(graph, neighbor, visited)
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

    visited := make(map[int]bool)
    fmt.Print("DFS Traversal: ")
    dfs(graph, 0, visited)  // Output: DFS Traversal: 0 1 3 4 2 5
}
