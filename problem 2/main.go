package main

import (
    "bufio"
    "fmt"
    "os"
)

// Function to check if a given string is a palindrome
func palindrome() {
    var test []string
    scanner := bufio.NewScanner(os.Stdin)

    // Read strings until a blank input is given
    fmt.Println("Enter strings to check if they are palindromes (press enter to stop):")
    for {
        if !scanner.Scan() {
            break
        }
        str := scanner.Text()
        if str == "" {
            break
        }
        test = append(test, str)
    }

    // Check each string if it is a palindrome
    for _, s := range test {
        if isPalindrome(s) {
            fmt.Printf("%s is a palindrome\n", s)
        } else {
            fmt.Printf("%s is not a palindrome\n", s)
        }
    }
}

// Helper function to check if a string is a palindrome
func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

func main() {
    palindrome()
}
