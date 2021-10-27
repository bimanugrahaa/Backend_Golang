package main

import "fmt"

//Get high * wide prime number
func primaSegiEmpat(high, wide, start int) {

	// your code here

	var isPrimeX int
	var prime []int
	var row int
	var col int
	var sum int

	//Check if start is 1 append some numbers
	if start == 1 {
		prime = append(prime, 2, 3, 5)
		start = 5
		isPrimeX = 3
	}

	// Check if start is even or odd number
	if start%2 == 0 {
		start += 1
	} else {
		start += 2
	}

	for i := start; isPrimeX < (high * wide); i += 2 {
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			isPrimeX++
			prime = append(prime, i)
		}
	}

	for _, value := range prime {
		if col <= wide {
			if row < high {
				fmt.Printf("%v\t", value)
				row++
				sum += value
			}
			if row == high {
				col++
				row = 0
				fmt.Println()
			}
		}
	}
	fmt.Printf("%v\n\n", sum)
}

func main() {

	primaSegiEmpat(2, 3, 13)

	/*

	 17 19

	 23 29

	 31 37

	 156

	*/

	primaSegiEmpat(5, 2, 1)

	/*

	 2  3  5  7 11

	 13 17 19 23 29

	 129

	*/

}
