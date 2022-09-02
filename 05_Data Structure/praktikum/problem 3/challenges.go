package main

import "fmt"

func PairSum(arr []int, target int) []int {
	temp := make(map[int]int)
	for i, number := range arr {
		if key, exist := temp[target-number]; exist {
			return []int{key, i}
		}
		temp[number] = i
	}
	return nil
}
func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
