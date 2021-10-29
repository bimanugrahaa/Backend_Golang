package main

import (
	"fmt"
	"math"
)

//Get the total minimum cost of frog to the last stone.
func Frog(jumps []int) int {

	// your code here

	step := 2                                             //Set the step from 2 because total cost of step 0 and 1 is known
	totalCost := make([]float64, len(jumps))              //Set the totalCost to the slice
	totalCost[0] = 0                                      //Total cost to the 1st stone is 0
	totalCost[1] = math.Abs(float64(jumps[1] - jumps[0])) //Total cost to the 2nd stone is h2 - h1

	//While frog step less than stones length.
	//Set the total cost of current stone.
	for step < len(jumps) {

		//Get the total cost of current stone by add the last total cost (or second last total cost)
		//with absolute difference value of current and last stone (or second last stone).
		lastStone := totalCost[step-1] + math.Abs(float64(jumps[step]-jumps[step-1]))
		twoLastStone := totalCost[step-2] + math.Abs(float64(jumps[step]-jumps[step-2]))

		//With math.Min find the minimum total cost from the last stone or second last stone.
		totalCost[step] = math.Min(lastStone, twoLastStone)
		step++
	}
	return int(totalCost[len(jumps)-1])
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
