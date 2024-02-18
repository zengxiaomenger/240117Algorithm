package main

func Zmggfunc(nums1 []int, m int, nums2 []int, n int) {
	ii := 0
	nums11 := make([]int, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums11[ii] = nums1[i]
			ii++
			i++
		} else {
			nums11[ii] = nums2[j]
			ii++
			j++
		}
		// fmt.Println(nums11)
	}
	// fmt.Println(i, j)
	if i < m {
		for i < m {
			nums11[ii] = nums1[i]
			ii++
			i++
		}
	} else {
		for j < n {
			nums11[ii] = nums2[j]
			ii++
			j++
		}
	}
	copy(nums1, nums11)
	// fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 3, 5, 7, 0, 0, 0}
	m := 4
	nums2 := []int{2, 4, 6}
	n := 3
	Zmggfunc(nums1, m, nums2, n)
}
