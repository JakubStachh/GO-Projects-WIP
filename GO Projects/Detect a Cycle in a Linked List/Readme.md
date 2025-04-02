# 📌 Detecting Cycle in a Linked List in Go

## 🚀 Description
This Go program detects whether a singly linked list contains a cycle. It uses **Floyd's Cycle Detection Algorithm**, also known as **Tortoise and Hare Algorithm**. The algorithm uses two pointers:

1. **Slow pointer** moves one step at a time.

2. **Fast pointer** moves two steps at a time.

If there is a cycle, both pointers will eventually meet at the same node. If no cycle exists, the fast pointer will eventually reach the end of the list (`nil`).

## 🔍 How It Works
- `hasCycle` function:

   - The function takes the head of the linked list as input.

   - Two pointers (`slow` and `fast`) are initialized to the head.

   - The slow pointer moves one step at a time, and the fast pointer moves two steps at a time.

   - If the fast pointer catches up with the slow pointer, a cycle is detected.

   - If the fast pointer reaches `nil`, the list does not contain a cycle.

## 🎯 Example Output
```sh
true
```
In this example:

- The program creates a linked list where the last node points back to the second node, creating a cycle.

- The program detects the cycle and prints `true`.

## 📂 Code Explanation
- `Node struct`: This struct defines a node in the linked list, with a `data` field for the value and a `next` field pointing to the next node.

- `hasCycle` function: This function detects a cycle in the linked list using two pointers (`slow` and `fast`).

- `main` function: This function creates a linked list with a cycle and calls `hasCycle` to check for the cycle.
