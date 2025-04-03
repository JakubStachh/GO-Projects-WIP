package main

import "fmt"

// Function to reverse a string
func reverseString(str string) string {
    runes := []rune(str)
    left, right := 0, len(runes)-1
    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }
    return string(runes)
}

func main() {
    str := "Hello, Go!"
    fmt.Println("Reversed String:", reverseString(str)) // Output: "!oG ,olleH"
}
