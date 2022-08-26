package main

import "fmt"

func main() {
	var number int
	var result string

	fmt.Print("Input : ")
	fmt.Scan(&number)

	if number > 1 {
		result = "Bilangan Prima"
		for i := 2; i < number; i++ {
			if number%i == 0 {
				result = "Bukan Bilangan Prima"
				break
			} else {
				result = "Bilangan Prima"
			}
		}
	} else {
		result = "Bukan Bilangan Prima"
	}
	fmt.Print(result)
}
