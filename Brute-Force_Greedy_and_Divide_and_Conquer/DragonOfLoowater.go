package main

import (
	"fmt"
	"sort"
)

//Get the minimum total knight height
func DragonOfLoowater(dragonHead, knightHeight []int) {

	// your code here

	var sumKnight int
	var sumDragon int

	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	if len(knightHeight) > len(dragonHead) {
		for _, valDragon := range dragonHead {
			for _, valKnight := range knightHeight {
				if valKnight >= valDragon {
					sumKnight += valKnight
					sumDragon++
					break
				}
			}
		}
	}

	if sumDragon < len(dragonHead) {
		fmt.Println("Knight fall")
	} else {
		fmt.Println(sumKnight)
	}
}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
