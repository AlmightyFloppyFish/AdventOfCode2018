package main

import (
	"fmt"
	io "io/ioutil"
	str "strings"
)

func main() {

	f, _ := io.ReadFile("input.txt")
	lines := str.Split(string(f), "\n")

	charsFull := make(map[int]int)

	for i := range lines {
		dupes := getDupes(lines[i])
		var (
			had2 bool
			had3 bool
		)
		for _, v := range dupes {
			if v == 2 && !had2 {
				charsFull[2]++
				had2 = true
			}
			if v == 3 && !had3 {
				charsFull[3]++
				had3 = true
			}
		}
	}
	fmt.Println(charsFull[2] * charsFull[3])
}

func getDupes(s string) map[rune]int {
	chars := make(map[rune]int, len(s))

	for _, v := range s {
		chars[rune(v)]++
	}
	return chars
}
