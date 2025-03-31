package main

import "fmt"

// Function to check if a string is a palindrome
func isPalindrome(str string) bool {
    left, right := 0, len(str)-1
    for left < right {
        if str[left] != str[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func main() {
    fmt.Println(isPalindrome("racecar"))  // Output: true
    fmt.Println(isPalindrome("hello"))    // Output: false
}
