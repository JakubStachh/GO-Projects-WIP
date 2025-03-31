package main

import "fmt"

// Define the Node struct
type Node struct {
    data int
    next *Node
}

// Function to detect cycle in a linked list
func hasCycle(head *Node) bool {
    slow, fast := head, head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if slow == fast {
            return true
        }
    }
    return false
}

func main() {
    // Creating a linked list with a cycle
    node1 := &Node{data: 1}
    node2 := &Node{data: 2}
    node3 := &Node{data: 3}
    node4 := &Node{data: 4}
    node5 := &Node{data: 5}

    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node5
    node5.next = node2 // Creating a cycle

    fmt.Println(hasCycle(node1)) // Output: true
}
