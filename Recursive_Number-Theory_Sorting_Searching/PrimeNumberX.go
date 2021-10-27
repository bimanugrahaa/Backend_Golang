package main

import "fmt"

//Get the "x th" Prime Number
func primeX(number int) int {

	// your code here

	isPrimeX := 3
	var i int

	if number == 1 {
		return 2
	}

	if number == 2 {
		return 3
	}

	if number == 3 {
		return 5
	}

	for i = 3; isPrimeX < number; i += 2 {

		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			isPrimeX++
		}
	}

	return i - 2
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
