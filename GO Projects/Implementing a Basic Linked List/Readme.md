# 📌 Singly Linked List Implementation in Go
 
## 🚀 Description
This Go program implements a **singly linked list** with basic operations:

- **Insert**: Adds a new node at the end of the list.

- **Display**: Prints the linked list in a readable format.

## 🔍 How It Works
- **Define** a `Node` **struct** with `data` (integer) and `next` (pointer to the next node).

- Implement an `insert` **function** to add nodes to the list.

- Implement a `display` **function** to traverse and print the list.

- Use these functions in `main()` to demonstrate functionality.

## 📜 Code Implementation
```sh
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
```
## 🎯 Example Output
```sh
10 -> 20 -> 30 -> nil  
```
## 📂 Key Features
- **Dynamic memory allocation** (Nodes are created dynamically)
- **Efficient insertion** (Appends to the end)
- **Traversal function** (`display`)
