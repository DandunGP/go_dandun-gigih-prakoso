package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var result = []int{}
	for i := 0; i < len(angka); i++ {
		var exist = 0
		for _, items := range angka {
			if string(angka[i]) == string(items) {
				exist += 1
			}
		}
		if exist == 1 {
			onlyOne, _ := strconv.Atoi(string(angka[i]))
			result = append(result, onlyOne)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
