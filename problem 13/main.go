//Write a program to find the intersection of two integer arrays
package main
import "fmt"
import "strings"
func intersection(arr1[]int,arr2[]int)[]int{
	elements :=make(map[int]bool)
	for _, num:=range arr1{
		elements[num]=true
	}
	var intersect []int
	for _, num :=range arr2{
		if elements[num]{
			intersect=append(intersect,num)
			elements[num]=false
		}
	}
return intersect

}
func readInts() []int {
    var input string
    fmt.Scan(&input)
    strs := strings.Split(input, ",")
    var nums []int
    for _, str := range strs {
        var num int
        fmt.Sscan(str, &num)
        nums = append(nums, num)
    }
    return nums
}

	func main() {
		fmt.Println("Enter elements of array1 (comma separated):")
		arr1 := readInts()
		fmt.Println("Enter elements of array2 (comma separated):")
		arr2 := readInts()
	
		result := intersection(arr1, arr2)
		fmt.Println("Intersect elements are:", result)
	}