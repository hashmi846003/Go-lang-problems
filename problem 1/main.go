//Write a program to find the sum of all elements in an integer array.
package main
import "fmt"
func main(){
	var  arr[10] int 
	for i:=0;i<10;i++{
		fmt.Scan(&arr[i])
	}
	sum:=0
	for i:=0;i<10;i++{
		sum=sum+arr[i]
	}
	fmt.Println("sum=",sum)


}

