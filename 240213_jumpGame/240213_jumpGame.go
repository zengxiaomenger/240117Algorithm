package main

import "fmt"

type node struct {
	index int
	val   int
	depth int
}

func countJump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] >= len(nums)-1 {
		return 1
	}
	depth := 0
	maxx := 0
	queue_num := 0
	var queue []node
	queue = append(queue, node{
		index: 0,
		val:   nums[0],
		depth: 0,
	})
	for len(queue) > 0 {
		//取出队头
		var node0 node = queue[0]
		//将队头可以跳到的所有位置入队，
		for i := node0.index + 1; i <= node0.index+node0.val; i++ {
			// fmt.Printf("i=%d\n", i)
			//防止重复入队
			if i < queue_num {
				continue
			}
			var node1 node = node{
				index: i,
				val:   nums[i],
				depth: node0.depth + 1,
			}
			queue = append(queue, node1)
			queue_num++
			//判断刚入队的这个是否能达到终点
			maxx1 := node1.index + node1.val
			if maxx1 >= len(nums)-1 {
				maxx = maxx1
				depth = node1.depth + 1
				break
			}
		}
		if maxx >= len(nums)-1 {
			break
		}
		queue = queue[1:]
	}
	return depth
}
func judgeJump(nums []int) bool {
	ans := true
	maxx := 0 //目前最远位置
	for i := 0; i < len(nums); i++ {
		if maxx < i {
			ans = false
			break
		}
		if maxx < i+nums[i] {
			maxx = i + nums[i]
		}
	}
	return ans
}
func main() {
	nums := []int{2, 3, 0, 1, 4}
	fmt.Println(countJump(nums))
	return
}
