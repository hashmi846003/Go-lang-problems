// Implement a function to convert a decimal number to binary.
package main

import "fmt"

func decimal(a int) []int {
    var b []int
    for a != 0 {
        b = append(b, a%2)
        a = a / 2
    }

    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return b
}

func main() {
    var a int
    fmt.Println("Enter the value:")
    fmt.Scan(&a)
    binary := decimal(a)
    fmt.Println("Binary representation:", binary)
}
