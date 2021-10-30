package main

import (
	"fmt"
	"strings"
	"time"
)

func LetterFrequency(str string) {
	alphabet := make(map[string]int)
	for _, val := range strings.ToLower(str) {
		if val >= 'a' && val <= 'z' {
			alphabet[string(val)]++
		}
	}
	fmt.Println(alphabet)
}

func main() {
	go LetterFrequency("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	time.Sleep(1 * time.Millisecond)
}
