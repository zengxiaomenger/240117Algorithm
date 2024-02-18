package main

import "fmt"

func rmEle(nums []int, val int) int {
	ii := 0 //已有的index
	for i, value := range nums {
		if value != val {
			nums[ii] = nums[i]
			ii++
		}
	}
	return ii
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(rmEle(nums, 3))
	fmt.Println("test")
	return
}
