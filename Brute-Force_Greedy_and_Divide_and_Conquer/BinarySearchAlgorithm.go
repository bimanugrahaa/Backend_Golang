package main

import "fmt"

//Search the position of x in given sorted array
func BinarySearch(array []int, x int) {

	// your code here

	beg := 0
	end := len(array) - 1
	mid := (beg + end) / 2

	for beg <= end {
		if array[mid] > x {
			end = mid - 1
		} else {
			beg = mid + 1
		}

		mid = (beg + end) / 2
	}

	if array[mid] != x {
		mid = -1
	}
	fmt.Println(mid)
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
