package main

import "fmt"

func primaSegiEmpat(high, wide, start int) {
	prime := primeX(start)
	j := 1
	a := 0
	total := 0
	for i := 0; i < high*wide; i++ {
		if j != high {
			fmt.Print(prime[a], " ")
			j++
		} else {
			fmt.Println(prime[a], " ")
			j = 1
		}
		total += prime[a]
		a++
	}
	fmt.Println(total)
}

func primeX(number int) []int {
	prime := []int{}

	for i := number + 1; i < 100; i++ {
		exist := false
		for j := 2; j < i/2+1; j++ {
			if i%j == 0 {
				exist = true
				break
			}
		}
		if !exist && i != 1 {
			prime = append(prime, i)
		}
	}

	return prime
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
