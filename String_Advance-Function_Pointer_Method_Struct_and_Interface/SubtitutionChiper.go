package main

import "fmt"

type student struct {
	name string

	nameEncode string

	score int
}

type Chiper interface {
	Encode() string

	Decode() string
}

func (s *student) Encode() string {

	var nameEncode = ""

	for _, val := range s.name {

		// encodeStr := val + 25 - (val-97)*2 		=> for ASCII <= 109
		// encodeStr := val - 1 - (val-110)*2		=> for ASCII > 109

		encodeStr := 219 - val
		nameEncode += string(encodeStr)
	}

	return nameEncode
}

func (s *student) Decode() string {

	var nameDecode = ""

	for _, val := range s.nameEncode {

		// decodeStr := val + 25 - (val-97)*2 		=> for ASCII <= 109
		// decodeStr := val - 1 - (val-110)*2		=> for ASCII > 109

		decodeStr := 219 - val
		nameDecode += string(decodeStr)
	}

	return nameDecode

}

func main() {

	var menu int

	var a = student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.nameEncode)

		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}

}
