package main

import (
	"fmt"
	"strconv"
)

//Get min and max number also their index based on given array
func FindMinAndMax(arr []int) string {

	// your code here

	min := 0
	max := 0
	minIndex := 0
	maxIndex := 0

	for key, val := range arr {
		if val < min {
			min = val
			minIndex = key
		}

		if val > max {
			max = val
			maxIndex = key
		}
	}

	return "min " + strconv.Itoa(min) + " index " + strconv.Itoa(minIndex) + " max " + strconv.Itoa(max) + " index " + strconv.Itoa(maxIndex)
}

func main() {

	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))

	// min: -2 index: 3 max: 8 index: 5

	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))

	// min: -5 index: 1 max: 22 index: 3

	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))

	// min: -21 index: 4 max: 9 index: 2

	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))

	// min: -1 index: 0 max: 18 index: 5

	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))

	// min: -20 index: 5 max: 7 index: 4

}
