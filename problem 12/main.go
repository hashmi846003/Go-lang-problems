//Implement a function to find the GCD (Greatest Common Divisor) of two integers
package main
import "fmt"
func gcd(a int,b int) int{
	for b!=0{
		a,b=b,a%b
	}
	return a

}
func main(){
	fmt.Println("Enter the two numbers to calculate the gcd")
	var c int
	var d int
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Print(gcd(c,d))
}