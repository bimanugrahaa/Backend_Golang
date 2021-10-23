package main

import (
	"fmt"
)

//Shift alphabet with specific offset
func caesar(offset int, input string) string {

	offsetMove := offset % 26 //Get the modulus to reduce the 'shift' time.

	var newStr string
	for _, value := range input {

		value += rune(offsetMove)

		//Get the alphabet back to 'a'
		if value > 122 {
			value -= 26
		}

		newStr += string(value)
	}

	return newStr
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))

	// bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// mnopqrstuvwxyzabcdefghijkl

}
