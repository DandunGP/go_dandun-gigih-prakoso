package main

import "fmt"

func pow(x, n uint) uint {
	var result uint = 1
	for n >= 1 {
		result *= x
		n--
	}
	return result
}

func main() {
	fmt.Println(pow(2, 63)) // 8
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
