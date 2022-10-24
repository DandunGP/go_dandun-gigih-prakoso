package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			if i*j >= 10 {
				fmt.Print(i*j, " ")
			} else {
				fmt.Print(i*j, "  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	cetakTablePerkalian(9)
}
