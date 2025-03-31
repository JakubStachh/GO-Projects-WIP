package main

import "fmt"

// Function to check if a number is prime
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(isPrime(7))  // Output: true
    fmt.Println(isPrime(10)) // Output: false
}
