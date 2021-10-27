package main

import (
	"fmt"
)

//Get the maximum sum from given array sequentially
func MaxSequence(arr []int) int {

	// your code here

	var min int
	var max int
	var minIndex int
	var maxIndex int
	var sum int
	var newArr []int

	for key, val := range arr {
		if val < 0 {
			if val <= min {
				max = min
				min = val
				maxIndex = minIndex
				minIndex = key
			}

			if val > max {
				max = val
				maxIndex = key
			}
		}
	}

	if minIndex < maxIndex {
		newArr = arr[minIndex+1 : maxIndex]
	} else {
		newArr = arr[maxIndex+1 : minIndex]
	}

	for _, val := range newArr {
		sum += val
	}

	return sum
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
