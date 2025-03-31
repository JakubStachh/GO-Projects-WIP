package main

import "fmt"

// Function to check if a string has unique characters
func hasUniqueCharacters(str string) bool {
    charMap := make(map[rune]bool)
    for _, char := range str {
        if charMap[char] {
            return false
        }
        charMap[char] = true
    }
    return true
}

func main() {
    fmt.Println(hasUniqueCharacters("abcde")) // Output: true
    fmt.Println(hasUniqueCharacters("aabbcc")) // Output: false
}
