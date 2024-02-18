package main

import "fmt"

func multiNum(nums []int) int {
	num := 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			num++
		} else {
			num--
			if num < 0 {
				pre = nums[i]
				num = 1
			}
		}
	}
	return pre
}
func main() {
	nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 4}
	fmt.Println(multiNum(nums))
	return
}
