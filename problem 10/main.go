//Implement a function to check if a given integer is a perfect square.
package main
import "fmt"
func main(){
	var a int
	fmt.Println("Enter the number for checking perfect square")
	fmt.Scan(&a)
	if sqrt(a) {
        fmt.Printf("%d is a perfect square.\n", a)
    } else {
        fmt.Printf("%d is not a perfect square.\n", a)
    }
}
func sqrt(a int)bool{
	//var a int
	//fmt.Scan(a)
	if a<0{
		return false
	}
	for i := 0; i*i <= a; i++ {
        if i*i == a {
            return true
        }
    }
    return false
		
			
		
	}
