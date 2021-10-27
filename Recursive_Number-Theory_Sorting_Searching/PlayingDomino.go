package main

import (
	"fmt"
)

//Get the suggested cards based on given deck
func playingDomino(cards [][]int, deck []int) interface{} {

	// your code here

	deckNow := 0
	sum := 0
	maxSum := 0
	var arrMax [2]int

	for i := 0; i < len(cards); i++ {
		if cards[i][0] == deck[deckNow] || cards[i][1] == deck[deckNow] {
			sum = cards[i][0] + cards[i][1]
			if sum > maxSum {
				maxSum = sum
				arrMax[0], arrMax[1] = cards[i][0], cards[i][1]
			}

		}
		deckNow++

		if cards[i][0] == deck[deckNow] || cards[i][1] == deck[deckNow] {
			sum = cards[i][0] + cards[i][1]
			if sum > maxSum {
				maxSum = sum
				arrMax[0], arrMax[1] = cards[i][0], cards[i][1]
			}

		}
		deckNow--

	}

	if maxSum != 0 {
		return arrMax
	} else {
		return "Tutup kartu"
	}
}

func main() {

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))

	// “tutup kartu”

}
