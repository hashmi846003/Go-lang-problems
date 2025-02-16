//Implement a function to reverse a string
package main
import "fmt"

func reverse(s string) string {
    runes := []rune(s)
    var i int = 0
    var j int = len(runes) - 1
    for i < j {
        runes[i], runes[j] = runes[j], runes[i]
        i++
        j--
    }
    return string(runes)
}

func main() {
    var str string
    fmt.Println("Enter the string")
    fmt.Scan(&str)
    fmt.Println(reverse(str))
}
