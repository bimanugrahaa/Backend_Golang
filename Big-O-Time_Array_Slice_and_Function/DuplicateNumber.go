package main

import (
	"fmt"
)

//Is the given string number has duplicate char or not.
func munculSekali(angka string) []int {

	check := make(map[rune]bool)
	res := []int{}

	for _, val := range angka {

		value := check[val]
		check[val] = !value
		if value {
			delete(check, val)
		}
	}

	for num, _ := range check {

		number := int(num - 48)
		res = append(res, number)
	}

	return res
}

func main() {

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}
