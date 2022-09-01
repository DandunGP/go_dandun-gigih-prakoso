package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var result = []int{}
	for i := 0; i < len(arr); i++ {
		exist := true
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				if i != j {
					result = append(result, i)
					result = append(result, j)
					exist = false
				}
			}
		}
		if !exist {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
