package main

import "fmt"

func playWithAsterisk(n int) {
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterisk(8)
}
