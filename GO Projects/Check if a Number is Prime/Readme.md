# 📌 Prime Number Checker in Go

## 🚀 Description
This Go program defines a function to check whether a given integer is a prime number. A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself.

## 🔍 How It Works
### The function `isPrime(n int) bool`:

- If `n` is less than or equal to 1, the function returns `false` because prime numbers are greater than 1.

- For numbers greater than 1, the function iterates from 2 to the square root of `n` (optimized check). If any number divides `n` without a remainder, `n` is not prime.

- If no divisors are found, the function returns `true`, indicating that `n` is a prime number.

## 🎯 Example Output
```sh
true
false
```
In this example:

- `isPrime(7)` returns `true` because 7 is a prime number.

- `isPrime(10)` returns `false` because 10 is divisible by 2 and 5, so it's not a prime number.

## 📂 Code Explanation
- `isPrime` Function:

   - The `isPrime` function efficiently checks if a number is prime by testing divisibility up to the square root of the number.

- `main` Function:

   - The `main` function calls `isPrime` with two test numbers (7 and 10) and prints the results.
