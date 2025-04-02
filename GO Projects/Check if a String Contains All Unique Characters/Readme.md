# 📌 Unique Characters Checker in Go

## 🚀 Description
This Go program checks whether a given string contains all unique characters. If the string has any repeating characters, it returns `false`; otherwise, it returns `true`.

## 🔍 How It Works
### The function `hasUniqueCharacters(str string) bool`:

- It iterates through each character in the string.

- It uses a `map[rune]bool` to track the characters that have already been encountered.

- If a character is encountered that is already in the map, it returns `false` immediately, indicating that the string has duplicates.

- If all characters are unique, the function returns `true`.

## 🎯 Example Output
```sh
true
false
```
In this example:

- `hasUniqueCharacters("abcde")` returns `true` because all characters in "abcde" are unique.

- `hasUniqueCharacters("aabbcc")` returns `false` because there are repeating characters.

## 📂 Code Explanation
- `hasUniqueCharacters` Function:

  - The function iterates over each character of the input string.

  - A map `charMap` keeps track of the characters that have been seen.

  - If a character is encountered twice, the function returns `false`. If no duplicates are found, it returns `true`.

- `main` Function:

  - The `main` function tests the `hasUniqueCharacters` function with two examples ("abcde" and "aabbcc") and prints the results.
