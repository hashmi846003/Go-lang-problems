package main

import "fmt"
func reversearray(a[]int)[]int{
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1{
		a[i],a[j]=a[j],a[i]
	}
	return (a)
}
func main(){
	fmt.Println(reversearray([]int{1,2,3,4,5}))
}
