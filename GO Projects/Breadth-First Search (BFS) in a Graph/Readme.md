# 📌 Breadth-First Search (BFS) in Go

## 🚀 Description
This Go program demonstrates how to perform a Breadth-First Search (BFS) traversal on an undirected graph represented as an adjacency list. BFS is used to explore nodes level by level starting from a given source node.

## 🔍 How It Works
The function `bfs(graph map[int][]int, start int)`:

- It initializes a `visited` map to track whether a node has been visited.

- A `queue` is initialized with the starting node and processed in a first-in, first-out (FIFO) manner.

- The algorithm then iterates over all the nodes in the graph, starting from the given `start` node:

   - It prints each node as it is visited.

   - For each visited node, it checks its neighbors. If a neighbor hasn't been visited, it is added to the queue.

- The BFS traversal proceeds until all reachable nodes from the `start` node have been visited.

## 🎯 Example Output
```sh
BFS Traversal: 0 1 2 3 4 5
```
In this example:

- The graph starts at node `0`, and the traversal explores all its neighbors (nodes `1` and `2`), then their respective neighbors, until all nodes are visited.

## 📂 Code Explanation
- `bfs` Function:

  - The `bfs` function uses a queue to maintain the order of nodes to be visited.

  - It ensures that each node is visited once by marking it in the `visited` map.

- `main` Function:

  - The `main` function defines a sample graph using an adjacency list representation.

  - It then calls the `bfs` function to start the BFS traversal from node `0`.
