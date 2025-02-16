//Write a program to find the Fibonacci sequence up to a specified number of terms.
package main
import "fmt"

func fibonacci(n int) []int {
    series := make([]int, n+1)
    if n >= 0 {
        series[0] = 0
    }
    if n >= 1 {
        series[1] = 1
    }
    for i := 2; i <= n; i++ {
        series[i] = series[i-1] + series[i-2]
    }
    return series
}

func main() {
    fmt.Println("Enter the value of n")
    var n int
    fmt.Scan(&n)
    series := fibonacci(n)
    for _, v := range series {
        fmt.Print(v, " ")
    }
    fmt.Println()
}

