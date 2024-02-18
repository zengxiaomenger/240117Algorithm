package main

import "fmt"

func rmDup(nums []int) int {
	ii := 1 //已有的index
	pre := nums[0]
	for i, value := range nums {
		if i == 0 {
			continue
		}
		if value != pre {
			nums[ii] = nums[i]
			ii++
			pre = nums[i]
		}
	}
	fmt.Println(ii)
	return ii
}
func rmDup2(nums []int) int {
	flag := 1 //当前元素数量
	ii := 1   //已有的index
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] == nums[i-1] {
			flag++
		} else {
			flag = 1
		}
		if flag <= 2 {
			nums[ii] = nums[i]
			ii++
		}
	}
	fmt.Println(ii)
	return ii
}
func main() {
	nums := []int{1, 2, 3, 3, 4, 4, 4, 4, 5}
	fmt.Println(rmDup2(nums))
	return
}
