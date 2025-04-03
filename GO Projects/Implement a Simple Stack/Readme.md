# 📌 Stack Implementation in Go

## 🚀 Description
This Go program implements a Stack using a slice. It supports basic stack operations:

- **Push**: Adds an element to the stack.

- **Pop**: Removes and returns the top element.

- **Peek**: Returns the top element without removing it.

## 🔍 How It Works
- **Define** a `Stack` struct that holds a slice of integers (`items`).

- Implement methods for `Push`, `Pop`, and `Peek`.

- Use the stack in `main()` to demonstrate its functionality.

## 📜 Code Implementation
```sh 
package main

import "fmt"

// Define a Stack structure
type Stack struct {
    items []int
}

// Push method to add an element to the stack
func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

// Pop method to remove an element from the stack
func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return -1
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return top
}

// Peek method to view the top element of the stack
func (s *Stack) Peek() int {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return -1
    }
    return s.items[len(s.items)-1]
}

func main() {
    stack := &Stack{}
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println("Top element:", stack.Peek()) // Output: 30
    fmt.Println("Popped element:", stack.Pop()) // Output: 30
    fmt.Println("Top element after pop:", stack.Peek()) // Output: 20
}
```
## 🎯 Example Output
```sh
Top element: 30  
Popped element: 30  
Top element after pop: 20  
```
## 📂 Key Features
- Dynamic size (Uses slice internally)
- Basic stack operations (`Push`, `Pop`, `Peek`)
- Handles empty stack cases
