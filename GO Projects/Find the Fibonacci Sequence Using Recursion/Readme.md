# 📌 Recursive Fibonacci in Go
## 🚀 Description
This Go program computes the n-th Fibonacci number using a recursive approach. The Fibonacci sequence is defined as:
##
$$
\large F(n)=F(n−1)+F(n−2)
$$
##
with the base cases:
##
$$
\large F(0)=0,F(1)=1
$$
##
<br></br>
## 🔍 How It Works
### The function `fibonacci(n int) int`:

- **Recursive Approach**: Calls itself with `n-1` and `n-2` until it reaches the base case.

- **Base Cases*: If `n` is `0` or `1`, it returns `n` directly.

- **Exponential Complexity**: Since it recalculates values multiple times, the time complexity is **O(2ⁿ)**.

The `main` function demonstrates how to compute the Fibonacci number for `n = 5` using recursion.

## 📜 Code Implementation
```sh
package main

import "fmt"

// Recursive function to get Fibonacci number
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println("Fibonacci of 5:", fibonacci(5)) // Output: 5
}
```

## 🎯 Example Output
```sh
Fibonacci of 5: 5
```
## 📂 Explanation
### `fibonacci(n int) int` Function:
- If `n` is `0` or `1`, it returns `n`.

- Otherwise, it calls itself for `n-1` and `n-2` and sums the results.

### `main` Function:
- Calls `fibonacci(5)`, which recursively computes the 5th Fibonacci number.

- Prints the result.
