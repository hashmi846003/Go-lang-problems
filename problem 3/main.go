package main

import (
    "fmt"
)

// Function to reverse a string
func reverseString(s string) string {
    // Convert string to rune slice to handle multi-byte characters
    runes := []rune(s)
    // Reverse the slice
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    // Convert back to string and return
    return string(runes)
}

func main() {
    str := "hello"
    reversedStr := reverseString(str)
    fmt.Println("Original string:", str)
    fmt.Println("Reversed string:", reversedStr)
}
