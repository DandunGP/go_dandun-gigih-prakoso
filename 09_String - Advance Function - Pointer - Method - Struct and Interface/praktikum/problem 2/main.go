package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	cipher := ""

	for _, inputs := range input {
		if strings.Contains(alphabet, string(inputs)) {
			cipher += string(alphabet[(strings.Index(alphabet, string(inputs))+offset)%(len(alphabet))])
		} else {
			cipher += string(inputs)
		}
	}
	return cipher
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
