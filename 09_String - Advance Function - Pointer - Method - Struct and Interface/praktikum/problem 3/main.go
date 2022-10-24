package main

import (
	"fmt"
)

func swap(a, b *int) {
	tempX := *a
	tempY := *b
	*a = tempY
	*b = tempX
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}
