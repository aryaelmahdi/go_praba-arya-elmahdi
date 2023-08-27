package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	input := s.name
	fmt.Println(input)
	for _, character := range input {
		offset := 9
		fmt.Println(character, offset)
		tmp := int(character) + offset
		if tmp > 122 {
			tmp -= 26
		}
		nameEncode += fmt.Sprintf("%c", tmp)
	}
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	input := s.name
	for _, character := range input {
		offset := -9
		tmp := int(character) + offset
		if tmp < 97 {
			tmp += 26
		}
		nameDecode += fmt.Sprintf("%c", tmp)
	}
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + "is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + "is : " + c.Decode())
	}
}
