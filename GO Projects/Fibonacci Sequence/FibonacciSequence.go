package main

import "fmt"

// Function to calculate Fibonacci using dynamic programming
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    fib := make([]int, n+1)
    fib[0], fib[1] = 0, 1

    for i := 2; i <= n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }

    return fib[n]
}

func main() {
    fmt.Println("Fibonacci of 10:", fibonacci(10)) // Output: 55
}
