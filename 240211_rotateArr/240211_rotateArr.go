package main

import "fmt"

func rotateArr0(nums []int) {
	s := 0
	t := len(nums) - 1
	for s < t {
		nums[s], nums[t] = nums[t], nums[s]
		s++
		t--
	}
}
func rotateArr(nums []int, k int) {
	k %= len(nums)
	rotateArr0(nums[:])
	rotateArr0(nums[0:k])
	rotateArr0(nums[k:len(nums)])
	fmt.Println(nums)
	return
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateArr(nums, 3)
	return
}
