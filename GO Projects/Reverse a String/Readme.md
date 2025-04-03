# 📌 String Reversal in Go

## 🚀 Description
This Go program reverses a given string. It takes an input string and returns its reverse order (e.g., `"Hello, Go!"` → `"!oG ,olleH"`).

## 🔍 How It Works
### The `reverseString()` function:

- Converts the string into a rune slice to handle multi-byte characters correctly.

- Uses two pointers (`left` and `right`) to swap characters from both ends.

- Returns the reversed string.

## 📂 Example
### Input:
```sh
str := "Hello, Go!"
fmt.Println("Reversed String:", reverseString(str))
```
### Processing:
1. Convert "Hello, Go!" → `['H', 'e', 'l', 'l', 'o', ',', ' ', 'G', 'o', '!']`.

2. Swap characters from start & end until the middle is reached.

3. Return reversed string `"!oG ,olleH"`.

### Output:
```sh
Reversed String: !oG ,olleH
```
## 🛠️ Implementation Details
Two-pointer technique:

Swaps characters from both ends, reducing operations.

Rune slice conversion:

Ensures support for UTF-8 characters (e.g., emojis, non-ASCII text).
