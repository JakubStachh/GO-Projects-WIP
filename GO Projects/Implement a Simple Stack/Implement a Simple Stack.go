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
