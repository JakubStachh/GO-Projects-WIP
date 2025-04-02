# 📌 Depth-First Search (DFS) in Go

## 🚀 Description
This Go program performs a Depth-First Search (DFS) traversal on a graph represented as an adjacency list. DFS is a graph traversal algorithm that explores as far as possible along each branch before backtracking.

## 🔍 How It Works
### The function `dfs(graph map[int][]int, node int, visited map[int]bool)`:

- The function begins at the given node and marks it as visited.

- It then recursively explores all unvisited neighbors of the current node.

- The traversal continues until all reachable nodes are visited.

### DFS Traversal in this program:

- The graph is represented using a map, where each key is a node, and the value is a slice of adjacent nodes (neighbors).

- A map `visited` keeps track of whether a node has been visited during the DFS.

## 🎯 Example Output
```sh
DFS Traversal: 0 1 3 4 2 5
```
In this example:

- The DFS starts at node `0`.

- The traversal follows the path `0 -> 1 -> 3 -> 4 -> 2 -> 5` as it explores all reachable nodes from `0`.

## 📂 Code Explanation
- `dfs` Function:

    - The `dfs` function is a recursive function that visits nodes in a depth-first manner.

    - It marks the current node as visited and then recursively visits each of its unvisited neighbors.

- `main` Function:

    - The `main` function defines the graph and initializes a `visited` map to keep track of visited nodes.

    - It calls the `dfs` function to start the traversal from node `0` and prints the result.
