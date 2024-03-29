package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var result float64
	for i := 0; i < len(s.score); i++ {
		result += float64(s.score[i])
	}
	result = result / float64(len(s.score))
	return result
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i := 0; i < len(s.score); i++ {
		if s.score[i] <= min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i := 0; i < len(s.score); i++ {
		if s.score[i] >= max {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input ", i, " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+"(", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+"(", scoreMin, ")")
}
