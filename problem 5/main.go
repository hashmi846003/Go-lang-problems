//Write a program to find the factorial of a given number.
package main
import "fmt"
func factorial(n int)int{
	//fact:=0
	if n==0{
		return 1
	}
	result:=1
	for i:=1;i<=n;i++{
		result*=i
	}
	return result
	
}
func main(){
	var n int
	fmt.Scanf("%d",&n)
	fmt.Println(factorial(n))
}