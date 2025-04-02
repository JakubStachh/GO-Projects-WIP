# 📌 Prime Numbers Finder in Go (Trial Division)

##🚀 Description
This Go program finds all prime numbers up to a given number `n` using the trial division method. The program consists of two key functions:

1. `isPrime`: Checks if a given number is prime by attempting to divide it by all numbers from `2` to the square root of the number.

2. `findPrimes`: Finds all prime numbers up to a given number `n` by iterating through all integers from `2` to `n` and checking each one with the `isPrime` function.

## 🔍 How It Works
### Trial Division Method:
The trial division method is a simple algorithm to check whether a number `n` is prime:

- A prime number is a natural number greater than 1 that cannot be formed by multiplying two smaller natural numbers.

- To check if a number `n` is prime, we check if it is divisible by any integer from `2` to the square root of `n`. If it is divisible, it is not a prime number.

### Program Flow:
1. `isPrime(n int) bool`: This function returns `true` if `n` is prime, otherwise returns `false`. It does this by checking all possible divisors from `2` to `sqrt(n)`.

2. `findPrimes(n int) []int`: This function returns a slice of integers containing all prime numbers up to `n`. It iterates through each number from `2` to `n` and checks if the number is prime using `isPrime`.

## 🎯 Example Output
```sh
Primes up to 50: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
```
For the input `n = 50`, the program outputs all the prime numbers up to `50`.

## 📝 Code Explanation
- `isPrime(n int) bool`:

   - This function checks whether the number `n` is prime.

   - It iterates from `2` to the square root of `n` and checks if n is divisible by any number. If it finds a divisor, the function returns `false`; otherwise, it returns `true`.

- `findPrimes(n int) []int`:

   - This function returns a slice of prime numbers from `2` to `n`.

   - It uses the `isPrime` function to check each number from `2` to `n` and adds the primes to the result slice.

- `main()`:

   - The `main` function calls `findPrimes(50)` to find all prime numbers up to `50` and prints the result.

## 🔧 Dependencies
### Go version: Go 1.16 or later is recommended to run the program.
