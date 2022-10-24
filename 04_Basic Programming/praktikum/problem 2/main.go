package main

import (
	"fmt"
)

func main() {
	var (
		nama         string
		studentScore int
	)

	fmt.Print("Masukan nama Siswa : ")
	fmt.Scan(&nama)
	fmt.Print("Masukan nilai Siswa : ")
	fmt.Scan(&studentScore)

	fmt.Println("--------------------------------")

	fmt.Println("Hasil Dari : ")
	fmt.Println("Nama : ", nama)
	fmt.Println("Nilai : ", studentScore)

	switch {
	case (studentScore >= 80 && studentScore <= 100):
		fmt.Println("Nilai A")
	case (studentScore >= 65 && studentScore < 80):
		fmt.Println("Nilai B")
	case (studentScore >= 50 && studentScore < 65):
		fmt.Println("Nilai C")
	case (studentScore >= 35 && studentScore < 49):
		fmt.Println("Nilai D")
	case (studentScore >= 0 && studentScore < 35):
		fmt.Println("Nilai E")
	default:
		fmt.Println("Nilai invalid")
	}
}

/*
Code 1

	func main() {
		var studentScore int = -1

		switch {
		case (studentScore >= 80 && studentScore <= 100):
			fmt.Println("A")
		case (studentScore >= 65 && studentScore < 80):
			fmt.Println("B")
		case (studentScore >= 50 && studentScore < 65):
			fmt.Println("C")
		case (studentScore >= 35 && studentScore < 49):
			fmt.Println("D")
		case (studentScore >= 0 && studentScore < 35):
			fmt.Println("E")
		default:
			fmt.Println("Nilai invalid")
		}
	}
*/
