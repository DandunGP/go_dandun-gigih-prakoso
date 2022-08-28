package main

import (
	"fmt"
)

func main() {
	var (
		T, r, luas float64
	)
	phi := 3.14

	fmt.Print("Tinggi : ")
	fmt.Scanf("%f\n", &T)
	fmt.Print("Jari-jari : ")
	fmt.Scanf("%f\n", &r)
	luas = 2*phi*r*r + 2*phi*r*T
	fmt.Print(luas)
}

/*
Code 1

func main() {
	T := 20.0
	r := 4.0
	phi := 3.14
	var luas float64

	luas = 2*phi*r*r + 2*phi*r*T

	fmt.Print(luas)
}
*/
