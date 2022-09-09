package main

import "fmt"

type Student struct {
	students map[string]int
}

func (s Student) Average() float64 {
	var result float64
	for _, value := range s.students {
		result += float64(value)
	}
	result = result / float64(len(s.students))
	return result
}

func (s Student) Min() (min int, name string) {
	for key, value := range s.students {
		if value <= min {
			min = value
			name = key
		}
	}
	return
}

func (s Student) Max() (max int, name string) {
	for key, value := range s.students {
		if value >= max {
			max = value
			name = key
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
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.students[name] = score
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+"(", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+"(", scoreMin, ")")
}
