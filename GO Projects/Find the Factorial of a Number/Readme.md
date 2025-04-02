📌 Factorial Calculation in Go (Recursive Approach)
🚀 Description
This Go program calculates the factorial of a given number using a recursive function. The factorial of a number n (denoted as n!) is the product of all positive integers from 1 to n. It is defined as:
##
$$
\large n!=n×(n−1)×(n−2)×...×1
$$
##
For example:

5! = 5 × 4 × 3 × 2 × 1 = 120

3! = 3 × 2 × 1 = 6

0! = 1 (by definition)

🔍 How It Works
This program implements a recursive function factorial(n int) int:

If n is 0, the function returns 1 (base case).

Otherwise, the function returns n * factorial(n-1), where the function calls itself with n-1 until it reaches 0.

Program Flow:
Base Case: When n == 0, return 1.

Recursive Case: When n > 0, return n * factorial(n-1).

Execution: The function keeps calling itself with n-1 until it reaches 0, at which point it starts returning values back up the call stack.

🎯 Example Output
sh
Copy
Edit
Factorial of 5: 120
For input n = 5, the program calculates:

5
!
=
5
×
4
×
3
×
2
×
1
=
120
5!=5×4×3×2×1=120
📝 Code Explanation
factorial(n int) int:

If n == 0, return 1.

Otherwise, return n * factorial(n-1).

Uses recursion to break down the problem into smaller subproblems.

main():

Calls factorial(5) to calculate the factorial of 5 and prints the result.
