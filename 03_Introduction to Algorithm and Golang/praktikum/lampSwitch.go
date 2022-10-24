package main

import "fmt"

func main() {
	var button int
	var lamp = "lampu mati"
	fmt.Print("Input : ")
	fmt.Scan(&button)

	for i := 1; i <= button; i++ {
		if button%i == 0 {
			if lamp == "lampu menyala" {
				lamp = "lampu mati"
			} else {
				lamp = "lampu menyala"
			}
		}
	}
	fmt.Println(lamp)
}
