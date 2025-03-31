package main

import "fmt"

// Define the Node struct
type Node struct {
    data int
    next *Node
}

// Insert a new node at the end
func insert(head *Node, data int) *Node {
    newNode := &Node{data: data}
    if head == nil {
        return newNode
    }
    temp := head
    for temp.next != nil {
        temp = temp.next
    }
    temp.next = newNode
    return head
}

// Display the linked list
func display(head *Node) {
    temp := head
    for temp != nil {
        fmt.Print(temp.data, " -> ")
        temp = temp.next
    }
    fmt.Println("nil")
}

func main() {
    var head *Node
    head = insert(head, 10)
    head = insert(head, 20)
    head = insert(head, 30)
    display(head) // Output: 10 -> 20 -> 30 -> nil
}
