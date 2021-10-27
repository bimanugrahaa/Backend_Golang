package main

import (
	"fmt"
)

//A B C Relations
// x + y + z = A
// xyz = B
// x^2 + y^2 + z^2 = C

//Get x, y, z value based on given A B C and their relations
func SimpleEquations(a, b, c int) {

	// your code here

	// x^2 + 2xy + 2xz + y^2 + 2yz + z^2
	// x^2 + y^2 + z^2 + 2xy + 2xz + 2yz = (x + y + z)^2
	// C + 2(xy + xz + yz) = A^2
	// (xy + xz + yz) = (A^2 - C)/2

	var x int
	var y int
	var z int
	var arrInt []int

	divAC := ((a * a) - c) / 2

out:
	for x = 1; x < divAC; x++ {
		for y = x + 1; y < divAC; y++ {
			for z = y + 1; z < divAC; z++ {
				if (x*y + x*z + y*z) == divAC {
					if x*y*z == b {
						arrInt = []int{x, y, z}
						break out
					}
				}
			}
		}
	}

	if len(arrInt) != 0 {
		fmt.Println(arrInt)
	} else {
		fmt.Println("No solution")
	}
}

func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

}
