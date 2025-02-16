package main

import "fmt"

func main() {
    var str string
    fmt.Println("Enter the string:")
    fmt.Scan(&str)

    vowels := 0
    consonants := 0

    for _, v := range str {
        if v == 'a' || v == 'A' || v == 'e' || v == 'E' || v == 'i' || v == 'I' || v == 'o' || v == 'O' || v == 'u' || v == 'U' {
            vowels++
        } else if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
            consonants++
        }
    }

    fmt.Println( vowels)
    fmt.Println( consonants)
}
