package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	for i := 0; i < len(arrayB); i++ {
		var exist = false
		for _, items := range arrayA {
			if arrayB[i] == items {
				exist = true
				break
			}
		}
		if !exist {
			arrayA = append(arrayA, arrayB[i])
		}
	}
	return arrayA
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
