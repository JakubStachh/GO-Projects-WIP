package main

import "fmt"

// Function to check if a number is prime using trial division
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

// Function to find all primes up to n using trial division
func findPrimes(n int) []int {
    var primes []int
    for i := 2; i <= n; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

func main() {
    primes := findPrimes(50)
    fmt.Println("Primes up to 50:", primes) // Output: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
}
