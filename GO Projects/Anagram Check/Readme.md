# 📌 Anagram Checker in Go

## 🚀 Description
This Go program checks whether two strings are anagrams of each other. Two strings are considered anagrams if they contain the same characters, with the same frequency, but potentially in a different order.

## 🔍 How It Works
The function `areAnagrams(str1, str2 string) bool`:

- It first checks if the lengths of the two strings are different. If so, they cannot be anagrams.

- It then counts the frequency of each character in the first string using a map.

- It then decreases the count for each character found in the second string.

- If any character has a negative count or is missing from the first string, the strings are not anagrams.

- If the strings have identical character frequencies, they are anagrams.

## 🎯 Example Output
```sh
true
false
```

- For `areAnagrams("listen", "silent")`, the output is `true` because "listen" and "silent" are anagrams.

- For `areAnagrams("hello", "world")`, the output is `false` because "hello" and "world" are not anagrams.

## 📂 Code Explanation
 
- `areAnagrams` Function:

   - The function uses a map (`count`) to track the frequency of each character in the first string.

   - It then checks if the second string matches the character counts from the first string.

- `main` Function:

   - The `main` function demonstrates calling the `areAnagrams` function with two pairs of strings, printing the results.
