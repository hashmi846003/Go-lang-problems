//Write a program to find the maximum and minimum elements in an integer array
package main
import "fmt"
func main(){
	var n int
fmt.Printf("Enter the value of elements in the array\n")
fmt.Scan(&n)
if n<=0{
	fmt.Println("Enter the value greater than zero")
	return
}
values:=make([]int, n)
fmt.Println("Values")
for i:=0;i<n;i++{
	fmt.Scan(&values[i])

}
max:=values[0]
for _,value:=range values{
	if value>max{
		max=value
	}
}
min:=values[0]
for _,value:=range values{
	if value<min{
		min=value
	}

}

fmt.Printf("The maximum value is: %d\n", max)
fmt.Printf("The minimum value is: %d\n", min)


}