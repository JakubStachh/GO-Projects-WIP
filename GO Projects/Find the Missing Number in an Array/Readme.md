# 📌 Find the Missing Number in Go

## 🚀 Description

This Go program finds the **missing number** in a sequence of integers ranging from `1` to `n`. The approach leverages the mathematical formula for the sum of the first `n` natural numbers and compares it with the actual sum of the given array to identify the missing number.

## 🔍 How It Works

### `findMissingNumber(arr []int, n int) -> int`:

- Computes the **expected sum** using the formula:

```sh
Sum = n * (n + 1) / 2
```
- Iterates through the given array to compute the **actual sum**.

- The missing number is determined as:
```sh
missing = expectedSum - actualSum
```
## 📜 Code Implementation
```sh
package main

import "fmt"

// Function to find missing number
func findMissingNumber(arr []int, n int) int {
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    for _, num := range arr {
        actualSum += num
    }
    return expectedSum - actualSum
}

func main() {
    arr := []int{1, 2, 4, 6, 3, 7, 8}
    n := 8
    fmt.Println("Missing number:", findMissingNumber(arr, n)) // Output: 5
}
```
## 🎯 Example Output

```sh
Missing number: 5
```
## 📂 Explanation

1. The input array contains numbers **from 1 to 8**, but one number is missing.

2. The function calculates the expected sum of numbers **from 1 to 8** (which is ``36``).

3. The actual sum of the given numbers is `31`.

4. The **missing number** is `36 - 31 = 5`.
