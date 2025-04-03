# 📌 Palindrome Checker in Go

## 🚀 Description
This Go program checks whether a given string is a palindrome. A palindrome is a word, phrase, or sequence that reads the same forward and backward (e.g., "racecar", "madam").

## 🔍 How It Works
### The `isPalindrome()` function:

   - Uses two pointers (`left` and `right`) to compare characters from both ends.

   - If a mismatch is found, it returns `false`.

   - If all characters match, it returns `true`.

## 📂 Example
### Input:
```sh
fmt.Println(isPalindrome("racecar"))
fmt.Println(isPalindrome("hello"))
```
### Processing:
1. `"racecar"` → Left & right match at every step → ✅ **Palindrome**

2. `"hello"` → 'h' ≠ 'o' → ❌ **Not a Palindrome**

### Output:
```sh
true
false
```

## 🛠️ Implementation Details
### **Two-pointer technique**:

- Compares characters from the start and end simultaneously.

### **Early exit optimization**:

- If a mismatch is found, it exits immediately.
