package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) interface{} {
	if strings.Contains(a, b) || strings.Contains(b, a) {
		if strings.Compare(a, b) == 1 {
			return b
		} else {
			return a
		}
	}
	return "Tidak Ada"
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
