package main

import (
	"fmt"
	"strings"
)

//Compare between two strings
func Compare(a, b string) string {

	if strings.Contains(a, b) || strings.Contains(b, a) {
		if strings.Compare(a, b) == 1 {
			return b
		} else if strings.Compare(a, b) == -1 {
			return a
		}
	}

	return "Not contains"
}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) // AKA

	fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	fmt.Println(Compare("KI", "KIJANG")) // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA")) // ILA

}
