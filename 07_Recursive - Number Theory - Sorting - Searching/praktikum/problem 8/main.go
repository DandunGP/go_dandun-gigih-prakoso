package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	key := make(map[string]int)

	var result []pair

	for _, value := range items {
		key[value]++
	}

	for item, total := range key {
		result = append(result, pair{item, total})
	}
	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}