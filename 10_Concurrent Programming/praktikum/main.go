package main

import (
	"fmt"
)

type Freq map[string]int

func Frequency(s string) Freq {
	m := Freq{}
	for _, r := range s {
		m[string(r)]++
	}
	return m
}

func main() {
	m := Freq{}
	texts := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	c := make(chan Freq, len(texts))
	for i := 0; i < len(texts); i++ {
		go func(text string) {
			c <- Frequency(text)
		}(string(texts[i]))
	}
	for _ = range texts {
		for r, v := range <-c {
			m[string(r)] += v
		}
	}
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
