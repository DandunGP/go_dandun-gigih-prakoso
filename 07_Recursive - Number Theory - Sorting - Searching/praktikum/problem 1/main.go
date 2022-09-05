package main

import "fmt"

func primeX(number int)int{
	prime := []int{}

	for i := 1; i < 100; i++{
		exist := 0
		for j:=2; j < i/2+1; j++{
			if i%j == 0{
				exist = 1
				break
			}
		}
		if (exist == 0 && i != 1){
			prime = append(prime, i)
		}
	}
	return prime[number-1]
}

func main(){
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}