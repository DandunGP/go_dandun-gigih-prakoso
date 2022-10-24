package main

import "fmt"

func primeNumber(number int) bool {
	result := true

	if number > 1 {
		for i := 2; i < number; i++ {
			if number%i == 0 {
				result = false
				break
			} else {
				result = true
			}
		}
	} else {
		result = false
	}

	return result
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
