# 📌 Fibonacci Sequence Calculation in Go (Dynamic Programming)

## 🚀 Description
This Go program calculates the `n`-th Fibonacci number using dynamic programming. 
It stores the intermediate Fibonacci values in an array (`fib`) to avoid redundant calculations, thus optimizing performance. 
The program uses an iterative approach to compute the Fibonacci sequence.

## 🔍 How It Works
The Fibonacci sequence is defined as:

- `F(0) = 0`

- `F(1) = 1`

- `F(n) = F(n-1) + F(n-2) for n >= 2`

####In this program, we use a dynamic programming approach:

1. We create an array `fib` where `fib[i]` will store the `i`-th Fibonacci number.

2. The first two Fibonacci numbers `fib[0]` and `fib[1]` are initialized to `0` and `1`, respectively.

3. We then iterate from `2` to `n`, and for each index `i`, we compute the Fibonacci number as the sum of the previous two Fibonacci numbers (`fib[i-1] + fib[i-2]`).

4. The result is stored in `fib[n]`.

## 🎯 Example Output
```sh
Fibonacci of 10: 55
```
For the input `n = 10`, the 10th Fibonacci number is `55`.

## 📂 Code Explanation
- `fibonacci(n int) int`:

    - This function calculates the `n`-th Fibonacci number using dynamic programming.

    - It uses an array `fib` to store intermediate Fibonacci numbers.

    - The function iterates through the array and fills it with Fibonacci values from `2` to `n`.

- `main()`:

    - The `main` function calls `fibonacci(10)` to compute the 10th Fibonacci number and prints the result.
