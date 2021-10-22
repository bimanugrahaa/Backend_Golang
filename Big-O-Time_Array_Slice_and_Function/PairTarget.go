package main

import "fmt"

//Get sum of a pair number equal to target
func PairSum(arr []int, target int) []int {

	res := make([]int, 0)
	keyJ := len(arr) - 1
	key := 0

	for key < keyJ {
		if arr[key]+arr[keyJ] == target {
			res = append(res, key, (keyJ))
			return res
		} else if arr[key]+arr[keyJ] < target {
			key++
		} else {
			keyJ--
		}
	}

	return res
}

func main() {

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}
