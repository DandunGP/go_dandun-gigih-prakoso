package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxSequence(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dMax, result := nums[0], nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		}
		dMax = max(0, dMax+num)
		result = max(dMax, result)
	}

	return result
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
