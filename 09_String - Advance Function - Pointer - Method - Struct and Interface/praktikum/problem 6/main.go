package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *Student) Encode() string {
	var nameEncode = ""
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	cipherAlpha := "zyxwvutsrqponmlkjihgfedcba"

	for _, items := range s.name {
		if strings.Contains(alphabet, string(items)) {
			nameEncode += string(cipherAlpha[strings.Index(alphabet, string(items))])
		}
	}
	return nameEncode
}

func (s *Student) Decode() string {
	var nameDecode = ""
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	cipherAlpha := "zyxwvutsrqponmlkjihgfedcba"

	for _, items := range s.nameEncode {
		if strings.Contains(cipherAlpha, string(items)) {
			nameDecode += string(alphabet[strings.Index(cipherAlpha, string(items))])
		}
	}
	return nameDecode
}

func main() {
	var menu int
	var a = Student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
