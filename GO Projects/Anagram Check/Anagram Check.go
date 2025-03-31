package main

import "fmt"

// Function to check if two strings are anagrams
func areAnagrams(str1, str2 string) bool {
    if len(str1) != len(str2) {
        return false
    }

    count := make(map[rune]int)
    for _, char := range str1 {
        count[char]++
    }

    for _, char := range str2 {
        count[char]--
        if count[char] < 0 {
            return false
        }
    }

    return true
}

func main() {
    fmt.Println(areAnagrams("listen", "silent"))  // Output: true
    fmt.Println(areAnagrams("hello", "world"))    // Output: false
}
