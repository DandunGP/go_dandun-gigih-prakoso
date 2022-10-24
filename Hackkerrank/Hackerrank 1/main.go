package main

import (
	"fmt"
	"strings"
)

func ReverseString(str string) string {
	char := "!@#$%^&*()-+=.,"
	strSps := strings.Split(str, " ")
	strNew := ""

	for i := len(strSps) - 1; i >= 0; i-- {
		strSplt := strings.Split(strSps[i], "")
		for i := len(strSplt) - 1; i >= 0; i-- {
			if strSplt[i] != " " {
				if strings.Contains(char, strSplt[i]) == false {
					if i != 0 {
						strNew += strSplt[i] + "_"
					} else {
						strNew += strSplt[i]
					}
				}
			}
		}
		strNew += " "
	}

	return strNew
}

func main() {
	fmt.Println(ReverseString("Hello World!@-+"))
	fmt.Println(ReverseString("I am a student"))
}
