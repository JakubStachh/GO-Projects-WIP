package main

import "fmt"

// Function to find primes up to n using Sieve of Eratosthenes
func sieveOfEratosthenes(n int) []int {
    primes := make([]bool, n+1)
    primes[0], primes[1] = false, false

    for i := 2; i*i <= n; i++ {
        if primes[i] == false {
            continue
        }
        for j := i * i; j <= n; j += i {
            primes[j] = false
        }
    }

    var result []int
    for i := 2; i <= n; i++ {
        if primes[i] {
            result = append(result, i)
        }
    }

    return result
}

func main() {
    primes := sieveOfEratosthenes(50)
    fmt.Println("Primes up to 50:", primes)  // Output: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
}
